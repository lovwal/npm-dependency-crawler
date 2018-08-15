package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/lovwal/npm-dependency-crawler/registry"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage())
		return
	}
	pkg := os.Args[1]
	r := registry.NewClient("https://skimdb.npmjs.com").Registry()

	d, err := r.GetDoc(pkg)
	if err != nil {
		fmt.Println("failure getting package", pkg, err)
		os.Exit(-1)
	}
	vers := ""
	if len(os.Args) > 2 {
		vers = os.Args[1]
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
	n.print("")
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
	for _, node := range n.Dependencies {
		node.print(indent + "  ")
	}
}

func usage() string {
	return `./npm-dependency-crawler <package> [version]`
}
