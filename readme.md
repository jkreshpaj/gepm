# GEPM [![Awesome](https://camo.githubusercontent.com/c9addde68ccb46540ce442b838a6a1617a5d7050/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f636f7665726167652d38302532352d79656c6c6f77677265656e2e7376673f6d61784167653d32353932303030)](https://github.com/flakaal/gepm)
> Go Express Package Manager


GEPM is a terminal browser, installer and manager for go packages.

  - Search for packages
  - Install a package by just typing the number
  - All packages dependencies are saved in ```packages.json``` so you can clone and run your projects on new environments

<p align="center">
  <img src="preview.gif"/>
</p>


### Installation

```sh
$ go get github.com/flakaal/gepm
```

### Examples

#### Search for package

```bash
$ gepm [package name]
```

Will return a list of all packages names, descriptions, authors matching ```[package name]``` . After typing the package number to install gepm will create a ```packages.json``` file on current directory which contains all the packages you installed with gepm.

#### Install packages from file
If you are cloning your code into another machine and dont have the required packages installed on it you can run ```gepm``` to install them from ```packages.json```

### Todos

 - Faster api search
 - gepm run to just add package to file.json

### Acknowledgement
 - <a href="https://github.com/daviddengcn/gcse">gcse</a> by <a href="https://github.com/daviddengcn">daviddengcn</a> for package search.
 - <a href="https://github.com/briandowns/spinner">spinner</a> by <a href="https://github.com/briandowns">briandowns</a>  for loading spinners.

License
----

MIT


**Free Software, Hell Yeah!**
