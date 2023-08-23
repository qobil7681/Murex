{{ if env "DOCGEN_TARGET=vuepress" }}
{{ if env "DOCGEN_TARGET=ignore-prefix" }}
### {{ end }}icon: arrow-down

---
{{ end }}<h1>Install Murex</h1>

{{ if env "DOCGEN_TARGET=" }}<h2>Table of Contents</h2>

<div id="toc">

- [Pre-Compiled Binaries (HTTPS download)](#pre-compiled-binaries-https-download)
- [Installing From A Package Manager](#installing-from-a-package-manager)
  - [ArchLinux](#archlinux)
  - [FreeBSD Ports](#freebsd-ports)
  - [Homebrew](#homebrew)
  - [MacPorts](#macports)
- [Compiling From Source](#compiling-from-source)
  - [Installation From Source Steps](#installation-from-source-steps)
- [Compiling On Windows](#compiling-on-windows)
- [Including Optional Builtins](#including-optional-builtins)
- [External Dependencies (Optional)](#external-dependencies-optional)
- [Recommended Terminal Typeface](#recommended-terminal-typeface)

</div>

{{ end }}## Supported Platforms

Linux, BSD and macOS are fully supported, with other platforms considered
experimental and/or having known limitations.

Windows is a supported platform however Murex doesn't aim to replace coreutils.
So, depending on your required use case, you may need additional 3rd party
software to provide those utilities.

There is a more detailed breakdown of known compatibility issues in the
[{{ if env "DOCGEN_TARGET=" }}docs/{{ end }}supported platforms]({{ if env "DOCGEN_TARGET=" }}docs{{ end }}/supported-platforms.md) document.

## Pre-Compiled Binaries (HTTPS download)

[![Version](version.svg)](DOWNLOAD.md)
[![CodeBuild](https://codebuild.eu-west-1.amazonaws.com/badges?uuid=eyJlbmNyeXB0ZWREYXRhIjoib3cxVnoyZUtBZU5wN1VUYUtKQTJUVmtmMHBJcUJXSUFWMXEyc2d3WWJldUdPTHh4QWQ1eFNRendpOUJHVnZ5UXBpMXpFVkVSb3k2UUhKL2xCY2JhVnhJPSIsIml2UGFyYW1ldGVyU3BlYyI6Im9QZ2dPS3ozdWFyWHIvbm8iLCJtYXRlcmlhbFNldFNlcmlhbCI6MX0%3D&branch=master)](DOWNLOAD.md)

If you wish to download a pre-compiled binary then head to the [DOWNLOAD](DOWNLOAD.md)
page to select your platform.

{{ if env "DOCGEN_TARGET=vuepress" }}
<!-- markdownlint-disable -->
<a href="DOWNLOAD.html" alt="download murex"><img src="/download.png" class="centre-image"/></a>
<!-- markdownlint-restore -->
{{ end }}

## Installing From A Package Manager

[![Packaging status](https://repology.org/badge/vertical-allrepos/murex.svg)](https://repology.org/project/murex/versions)

### ArchLinux

From AUR: [https://aur.archlinux.org/packages/murex(https://aur.archlinux.org/packages/murex)

    wget -O PKGBUILD 'https://aur.archlinux.org/cgit/aur.git/plain/PKGBUILD?h=murex'
    makepkg --syncdeps --install 


### FreeBSD Ports

Murex is available in [FreeBSD ports](https://www.freebsd.org/ports/).

### Homebrew

    brew install murex
    
### MacPorts

    sudo port install murex

## Compiling From Source

[![Tests](https://github.com/lmorg/murex/actions/workflows/go-tests.yaml/badge.svg?branch=master)](https://github.com/lmorg/murex/actions/workflows/go-tests.yaml)

**Prerequisites:**

You will need `go` (Golang) compiler, `gcc` (C compiler) and `git` installed
as well as your `$GOPATH` environmental variable set. 

> Go 1.18 or higher is required.

These should be easy to install on most operating systems however Windows is a
lot more tricky with regards to `gcc`. Please check with your operating systems
package manager first but see further reading below if you get stuck.

**Further Reading:**

* [How to install Go](https://golang.org/doc/install)
* [How to install git](https://github.com/git-guides/install-git)
* [How to install gcc](https://gcc.gnu.org/install/)
* [How to set GOPATH](https://github.com/golang/go/wiki/SettingGOPATH)

### Installation From Source Steps

The following instructions are assuming you're compiling on a POSIX-compatible
system like Linux, BSD or macOS. Compiling from source is untested on Plan 9
(if you run into issues there then please use the pre-compiled binary for that
platform) and Windows. In the case of Windows you may run into issues with the
`gcc` installation and some of the commands below will need to be adapted (eg
`murex.exe` used instead of `./murex`).

> Compiling from source is not recommended unless you already have a strong
> understanding of compiling Go projects for your specific platform.

**Importing the source code:**

At present, Murex depends on being in a specific directory hierarchy for
the tests to work and packages to import correctly. These instructions will
talk you through creating that initial structure ready to import the source
into. Experienced users in Go may opt to ignore some of these steps and run
`go get -u github.com/lmorg/murex` instead. While this _should_ work in most
cases, it is difficult to run automated tests to ensure any updates doesn't
break the `go get` import tool. And thus that approach is not officially
supported. If you are in any doubt, please follow the `git clone` process
below.

First create the directory path and clone the source into the appropriate
directory structure.

    mkdir -p $GOPATH/src/github.com/lmorg/murex
    cd $GOPATH/src/github.com/lmorg/murex
    git clone https://github.com/lmorg/murex .

At this point you can add and remove any optional builtins by following the
instructions on this located further down this document. This is entirely
optional as Murex attempts to ship with sane defaults.

**Compile the code:**

    go build github.com/lmorg/murex

**Test the executable (optional stage):**

    ./murex -c 'g behavioural/*.mx -> foreach f { source $f }; test run *

**Start the shell:**

    ./murex

## Compiling On Windows

For building on Windows 10, a GCC compiler is needed for the [go-sqlite3]( https://github.com/mattn/go-sqlite3)
module.

Build instructions for the module support building Murex as well: https://github.com/mattn/go-sqlite3#windows

1. Install a Windows `gcc` toolchain, for example https://jmeubank.github.io/tdm-gcc
2. Ensure that `gcc` appears in the Windows path
3. `go build` within the Murex source path

These instructions are expected to work on other versions of Windows but, at
time of writing, haven't yet been officially tested.

## Including Optional Builtins

Some optional builtins will be included by default, however there may be others
you wish to include which are not part of the default build (such as `qr`). To
add them, copy (or symlink) the applicable include file from
`builtins/import_src` to `builtins/import_build`.

A tool will be introduced in a later version to automate this.

## External Dependencies (Optional)

Some of Murex's extended features will have additional external dependencies.

* `aspell`: This is used for spellchecking. Murex will automatically enable or
  disable spellchecking based on whether `aspell` can be found in your `$PATH`.
  [http://aspell.net](http://aspell.net)

## Recommended Terminal Typeface

This is obviously just a subjective matter and everyone will have their own
personal preference. However if I was asked what my preference was then that
would be [Hasklig](https://github.com/i-tu/Hasklig). It's a clean typeface
based off Source Code Pro but with a few added ligatures - albeit subtle ones
designed to make Haskell more readable. Those ligatures also suite Murex
pretty well. So the overall experience is a clean and readable terminal.