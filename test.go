package main

import (
    "fmt"
    "github.com/flosch/pongo2"
)

func main() {
    // Compile the template first (i. e. creating the AST)
    tpl, err := pongo2.FromString("Hello from a {{ lang|capfirst }} script!")
    if err != nil {
        panic(err)
    }
    // Now you can render the template with the given
    // pongo2.Context how often you want to.
    out, err := tpl.Execute(pongo2.Context{"lang": "go"})
    if err != nil {
        panic(err)
    }
    fmt.Println(out) // Output: Hello Florian!
}