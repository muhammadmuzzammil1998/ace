#

![ace](ace.png)

ace is a resource update notifier which triggers a webhook or a script when resources are updated.

## Build

Requires [git](https://git-scm.com/download/win) to clone and [Go](https://golang.org/dl/) to build.

Clone and navigate to directory

```sh
$ git clone https://github.com/muhammadmuzzammil1998/ace.git
$ cd ace
```

Get dependencies and build

```sh
$ go get -d .
$ go build
```

## Installation

### From releases

Follow the guide on the [releases](https://github.com/muhammadmuzzammil1998/ace/releases) page for detailed instructions.

#### Linux (.deb)

Download `.deb` file for ace from the [releases](https://github.com/muhammadmuzzammil1998/ace/releases) page.

```sh
$ wget https://github.com/muhammadmuzzammil1998/ace/releases/download/{version}/ace-linux-{amd64/386}.deb
```

Install using `dpkg`

```sh
$ sudo dpkg -i ace-linux-{amd64/386}.deb
```

#### Windows

Start PowerShell as an admin

Download `.exe` file for ace from the [releases](https://github.com/muhammadmuzzammil1998/ace/releases) page.

```ps
PS > Invoke-WebRequest https://github.com/muhammadmuzzammil1998/ace/releases/download/{version}/ace-windows-{amd64/386}.exe -OutFile ace.exe
```

Move `ace.exe` to `C:\Windows` or store it somewhere and add it to path.

```ps
PS > mv .\ace.exe C:\Windows\ace.exe
```

#### Other OSs

You can download the [binary already built](https://github.com/muhammadmuzzammil1998/ace/releases) for your system or [build it yourself](https://github.com/muhammadmuzzammil1998/ace#build). To install, see [install from source](https://github.com/muhammadmuzzammil1998/ace#from-source)

#### NOTE: This should be obvious but still

- Adapt `{version}` number. Check version number from [here](https://github.com/muhammadmuzzammil1998/ace/releases).
- Choose your architecture, `amd64` for 64 bit and `386` for 32 bit systems.

### From source

Clone and navigate to directory

```sh
$ git clone https://github.com/muhammadmuzzammil1998/ace.git
$ cd ace
```

Get dependencies and build

```sh
$ go get -d .
$ go install
```

### Via `go get`

```sh
$ go get -t -v -u muzzammil.xyz/ace
Fetching https://muzzammil.xyz/ace?go-get=1
Parsing meta tags from https://muzzammil.xyz/ace?go-get=1 (status code 200)
get "muzzammil.xyz/ace": found meta tag get.metaImport{Prefix:"muzzammil.xyz/ace", VCS:"git", RepoRoot:"https://github.com/muhammadmuzzammil1998/ace"} at https://muzzammil.xyz/ace?go-get=1
muzzammil.xyz/ace (download)
Fetching https://muzzammil.xyz/pagehashgo?go-get=1
Parsing meta tags from https://muzzammil.xyz/pagehashgo?go-get=1 (status code 200)
get "muzzammil.xyz/pagehashgo": found meta tag get.metaImport{Prefix:"muzzammil.xyz/pagehashgo", VCS:"git", RepoRoot:"https://github.com/muhammadmuzzammil1998/pagehash-go"} at https://muzzammil.xyz/pagehashgo?go-get=1
muzzammil.xyz/pagehashgo (download)
muzzammil.xyz/ace
$ ace -version
ace v1.19.7.1
  by Muhammad Muzzammil.
  https://muzzammil.xyz
```

## [Documentation](DOCUMENTATION.md)

## Contributions

Contributions are welcome but kindly follow the [Code of Conduct](CODE_OF_CONDUCT.md) and guidlines. Please don't make Pull Requests for typographical errors, grammatical mistakes, "sane way" of doing it, etc. Open an issue for it. Thanks!
