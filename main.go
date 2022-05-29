package main

import (
	"flag"
)

func main() {
	// define arguements
	dir := flag.String("dir", ".", "directory")
	mode := flag.Int("mode", 0, "File (0) or directory (1)")
	expr := flag.String("expr", "^$", "Regular expression to extract data")
	name_template := flag.String("name-template", "", "Name template")
	dry_run := flag.Bool("dry-run", false, "Dry run only")
	flag.Parse()

	app(*dir, *mode, *expr, *name_template, *dry_run)
}
