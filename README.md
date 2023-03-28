#⚡️HyperCloud CLI

The HyperCloud CLI is the official command line tool of the HyperCloud Network.

### Running
`hyper` provides boilerplate code and packaging for web applications. It's fully ESM aware using
[vite](https://vite.dev) and *typescript* out of the box as well as support for *sass* and *less*
provided the packages are installed. `hyper` also provides a wrapper around package installation.

Command Line Options:
```

██╗  ██╗██╗   ██╗██████╗ ███████╗██████╗
██║  ██║╚██╗ ██╔╝██╔══██╗██╔════╝██╔══██╗
███████║ ╚████╔╝ ██████╔╝█████╗  ██████╔╝
██╔══██║  ╚██╔╝  ██╔═══╝ ██╔══╝  ██╔══██╗
██║  ██║   ██║   ██║     ███████╗██║  ██║
╚═╝  ╚═╝   ╚═╝   ╚═╝     ╚══════╝╚═╝  ╚═╝
(c) 2020-2023 HyperCloud.network



Usage: hyper [command]
-----------------------------
        init [dir]      Inits a new project
        build           Builds the current project.
        install [args]  Installs dependencies
        serve           Serves the current project for development.
        help            this help message.
```
* init [dir] will initialize a new project at the directory specified.
* build will build the project in the current directory.
* install will install any project dependencies. You can add them by giving it arguments.
* serve will provide a HMR-capable dev server for quick prototyping.
* help is redundant at this point. 
### Build

To build, run `make`
To build for a different platform than the one you are running on...
```
export GOOS=linux     # or windows or darwin
export GOARCH=amd64   # or arm64 or IA32, please don't use IA32...
make build
```
and go look in the `bin` folder for the executable.
### Install

To install, run `make install`
On linux machines it's typical to put the binary in `/usr/bin/hyper` but
I'll leave that up to you, for now it installs in *go*'s bin at
`%HOME%/go/bin` which you should have on your path if you develop in go.

If not, you can place it in `%HOME%/bin` for just your user and add that
to your path with `export PATH=$PATH:/home/me/bin` where `/home/me` is your 
`%HOME%` for your OS. Yours **_will_** be different.

### Package

Coming Soon...