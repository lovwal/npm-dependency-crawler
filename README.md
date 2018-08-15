### npm-dependency-crawler
A silly command line tool to view all dependencies and their dependencies for a specific npm package.

### Usage
```
./npm-dependency-crawler <package> [version]
```

### Example
```
$ ./npm-dependency-crawler react
Crawling react at version 16.4.2
Crawl completed in 3.910463353s
react 16.4.2
  fbjs 0.8.16
    core-js 1.0.0
    isomorphic-fetch 2.1.1
      whatwg-fetch 0.8.2
      node-fetch 1.0.1
        encoding 0.1.11
          iconv-lite ~0.4.4
    loose-envify 1.0.0
      js-tokens 1.0.1
    object-assign 4.1.0
    promise 7.1.1
      asap ~2.0.3
    setimmediate 1.0.5
    ua-parser-js 0.7.9
  loose-envify 1.1.0
    js-tokens 1.0.1
  object-assign 4.1.1
  prop-types 15.6.0
    fbjs 0.8.16
      core-js 1.0.0
      isomorphic-fetch 2.1.1
        node-fetch 1.0.1
          encoding 0.1.11
            iconv-lite ~0.4.4
        whatwg-fetch 0.8.2
      loose-envify 1.0.0
        js-tokens 1.0.1
      object-assign 4.1.0
      promise 7.1.1
        asap ~2.0.3
      setimmediate 1.0.5
      ua-parser-js 0.7.9
    loose-envify 1.3.1
      js-tokens 3.0.0
    object-assign 4.1.1
Total number of dependencies for react 35
