# fsgo

<!-- Build Status, is a great thing to have at the top of your repository, it shows that you take your CI/CD as first class citizens -->
<!-- [![Build Status](https://travis-ci.org/jjasghar/ibm-cloud-cli.svg?branch=master)](https://travis-ci.org/jjasghar/ibm-cloud-cli) -->

## Scope

A FileSystem Abstraction System for Go, based on the package [Afero](https://https://github.com/spf13/afero). This project is a subset of the functionality of Fsgo that eliminates the use of the crypto dependency. The Fsgo project includes concrete implementations of the abstraction that support multiple filesystems that are not needed for unit testing in a secure environment.

Fsgo is a filesystem framework providing a simple, uniform and universal API
interacting with any filesystem, as an abstraction layer providing interfaces,
types and methods. Fsgo has an exceptionally clean interface and simple design
without needless constructors or initialization methods.

Fsgo is also a library providing a base set of interoperable backend
filesystems that make it easy to work with fsgo while retaining all the power
and benefit of the os and ioutil packages.

Fsgo provides significant improvements over using the os package alone, most
notably the ability to create mock and testing filesystems without relying on the disk.

It is suitable for use in any situation where you would consider using the OS
package as it provides an additional abstraction that makes it easy to use a
memory backed file system during testing. It also adds support for the http
filesystem for full interoperability.

## Fsgo Features

* A single consistent API for accessing a variety of filesystems
* Interoperation between a variety of file system types
* A set of interfaces to encourage and enforce interoperability between backends
* An atomic cross platform memory backed file system
* Support for compositional (union) file systems by combining multiple file systems acting as one
* Specialized backends which modify existing filesystems (Read Only, Regexp filtered)
* A set of utility functions ported from io, ioutil & hugo to be fsgo aware
* Wrapper for go 1.16 filesystem abstraction `io/fs.FS`

# Using Fsgo

Fsgo is easy to use and easier to adopt.

A few different ways you could use Fsgo:

* Use the interfaces alone to define your own file system.
* Wrapper for the OS packages.
* Define different filesystems for different parts of your application.
* Use Fsgo for mock filesystems while testing

## Step 1: Install Fsgo

First use go get to install the latest version of the library.

    $ go get github.com/IBM/fsgo

Next include Fsgo in your application.
```go
import "github.com/IBM/fsgo"
```

## Step 2: Declare a backend

First define a package variable and set it to a pointer to a filesystem.
```go
var AppFs = fsgo.NewMemMapFs()

or

var AppFs = fsgo.NewOsFs()
```
It is important to note that if you repeat the composite literal you
will be using a completely new and isolated filesystem. In the case of
OsFs it will still use the same underlying filesystem but will reduce
the ability to drop in other filesystems as desired.

## Step 3: Use it like you would the OS package

Throughout your application use any function and method like you normally
would.

So if my application before had:
```go
os.Open("/tmp/foo")
```
We would replace it with:
```go
AppFs.Open("/tmp/foo")
```

`AppFs` being the variable we defined above.

## Notes

**NOTE: This repository has been configured with the [DCO bot](https://github.com/probot/dco).
When you set up a new repository that uses the Apache license, you should
use the DCO to manage contributions. The DCO bot will help enforce that.
Please contact one of the IBM GH Org stewards.**

<!-- Questions can be useful but optional, this gives you a place to say, "This is how to contact this project maintainers or create PRs -->
If you have any questions or issues you can create a new [issue here][issues].

Pull requests are very welcome! Make sure your patches are well tested.
Ideally create a topic branch for every separate change you make. For
example:

1. Fork the repo
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Added some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

## License

All source files must include a Copyright and License header. The SPDX license header is 
preferred because it can be easily scanned.

If you would like to see the detailed LICENSE click [here](LICENSE).

```text
#
# Copyright 2020- IBM Inc. All rights reserved
# SPDX-License-Identifier: Apache2.0
#
```
