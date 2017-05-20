## Satisfy

> Satisfy is a CLI tool that will allow the user to satisfy an interface/s of their choice
by generating code that implements its methods.

_**Special thanks to [Brian Downs](https://github.com/briandowns) - "satisfy" is heavily influenced by his tool, [TODO-VIEW](https://github.com/briandowns/todo-view) in terms of design.**_ 

### Basic usage

**Searching interfaces**

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

**Generate method signatures**

_FYI: **satisfy** will display an error message and will exit if the interface name provided is not found in the current tree._

```sh
$ satisfy implement <interface-name> [<option>] <type>, [<option> <type>...]
```

```sh
$ satisfy implement CoolInterface CoolStruct
```

Output:
```go
func (c CoolStruct) punch() {

}

func (c CoolStruct) kick() {

}
```

You can specify whether or not you want the receivers to be pointer types or value types.

* -p pointer
* -v value

You can switch between the two flags within the same command, so every type following each flag will have that receiver type in the signature.

```sh
$ satisfy implement KarateChopper -p Struct1 Struct2 -v struct3 -p struct4
```

Output:
```go
func (c *Struct1) punch(s string) (string, error) {

}

func (c *Struct1) kick() {

}

func (c *Struct2) punch(s string) (string, error) {

}


func (c *Struct2) kick() {

}

func (c struct3) punch(s string) (string, error) {

}


func (c struct3) kick() {

}

func (c *struct4) punch(s string) (string, error) {

}


func (c *struct4) kick() {

}
```

### Contributions
Contributions are welcome!
If you have any ideas regarding additional functionality, or want to improve something, please fork and open a pull request.  That would be awesome!