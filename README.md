### Satisfy

> Satisfy is a CLI tool that will allow the user to satisfy an interface/s of their choice
by generating code that implements its methods.

#### Basic usage

Search all interfaces in current tree and display their names, methods, and containing file
```sh
$ satisfy isearch

Interface Name: IFacer2 - commands/sampler2.go
        Do2(n int, s []string) error
        While2(s string) (string, error)
Interface Name: sampler - sampler.go
        punch()
        kick()
```

Search interfaces in current tree and only display exported interfaces
```sh
$ satisfy isearch filter -e

Interface Name: IFacer2 - commands/sampler2.go
```

Search all interfaces but do not display methods
```sh
$ satisfy isearch filter

Interface Name: IFacer2 - commands/sampler2.go
Interface Name: sampler2 - commands/sampler2.go
Interface Name: IFacer - sampler.go
Interface Name: sampler - sampler.go
```