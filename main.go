package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/lovwal/npm-dependency-crawler/registry"
)

func main() {
	var outputJson bool
	flag.BoolVar(&outputJson, "json", false, "print output in json")
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println(usage())
		return
	}
	pkg := args[0]
	r := registry.NewClient("https://skimdb.npmjs.com").Registry()

	d, err := r.GetDoc(pkg)
	if err != nil {
		fmt.Println("failure getting package", pkg, err)
		os.Exit(-1)
	}
	vers := ""
	if len(args) > 1 {
		vers = args[1]
	} else {
		vers = d.DistTags["latest"]
	}
	fmt.Println("Crawling", pkg, "at version", vers)
	start := time.Now()
	n := &Node{
		Name:         d.Name,
		Version:      vers,
		Dependencies: make(map[string]*Node),
	}
	dependencies := n.getDependencies(r, d)
	fmt.Println("Crawl completed in", time.Since(start))
	if outputJson {
		j, err := json.MarshalIndent(n, "", "\t")
		if err != nil {
			fmt.Println("failed to convert result to json:", err)
		} else {
			fmt.Println(string(j))
		}
	} else {
		n.print("")
	}
	fmt.Println("Total number of dependencies for", n.Name, dependencies)
}

type Node struct {
	Name         string
	Version      string
	Dependencies map[string]*Node
}

func (n *Node) getDependencies(r *registry.Registry, d *registry.Doc) int {
	dependecies := len(d.Versions[n.Version].Dependencies)
	for id, vers := range d.Versions[n.Version].Dependencies {
		doc, err := r.GetDoc(id)
		if err != nil {
			fmt.Println("aborting crawl on submodule", id, "due to error", err)
			return dependecies
		}
		node := &Node{
			Name:         id,
			Version:      strings.TrimPrefix(vers, "^"),
			Dependencies: make(map[string]*Node),
		}
		n.Dependencies[id] = node
		dependecies += node.getDependencies(r, doc)
	}
	return dependecies
}

func (n *Node) print(indent string) {
	fmt.Println(indent+n.Name, n.Version)
	var pkgs []string
	for _, node := range n.Dependencies {
		pkgs = append(pkgs, node.Name)
	}
	sort.Strings(pkgs)
	for _, pkg := range pkgs {
		n.Dependencies[pkg].print(indent + "  ")
	}
}

func usage() string {
	return `./npm-dependency-crawler [flags] <package> [version]
	-json output result in json format`
}
