package main

import (
	"github.com/mgutz/goa"
	f "github.com/mgutz/goa/filter"
	. "github.com/mgutz/godo"
	"github.com/mgutz/str"
)

// Project is the local project.
func Tasks(p *Project) {
	p.Task("dist", []string{"lint", "readme"})

	p.Task("lint", func() {
		Run("golint .")
		Run("gofmt -w -s .")
		Run("go vet .")
	})

	p.Task("readme", func() {
		Run("godocdown -o README.md")
		// add godoc
		goa.Pipe(
			f.Load("./README.md"),
			f.Str(str.ReplaceF("--", "\n[godoc](https://godoc.org/github.com/mgutz/goa)\n", 1)),
			f.Write(),
		)
	})
}

func main() {
	Godo(Tasks)
}