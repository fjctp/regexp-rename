package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	// define arguements
	dir := flag.String("dir", ".", "directory")
	mode := flag.Int("mode", 0, "File (0) or directory (1)")
	expr := flag.String("expr", "^$", "Regular expression to extract data")
	name_template := flag.String("name-template", "", "Name template")
	dry_run := flag.Bool("dry-run", false, "Dry run only")
	flag.Parse()

	// get abs path
	dir_abs, err := filepath.Abs(*dir)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Work dir:", dir_abs)

	// loop through dir
	files, err := ioutil.ReadDir(dir_abs)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() && (*mode == 1) {
			var re = regexp.MustCompile(*expr)

			old_name := file.Name()
			new_name := re.ReplaceAllString(old_name, *name_template)

			is_skip := old_name == new_name
			if is_skip {
				log.Print(old_name, " -> ", "Skip")
			} else {
				log.Print(old_name, " -> ", new_name)
			}

			if !(*dry_run) && !is_skip {
				err := os.Rename(
					filepath.Join(dir_abs, old_name),
					filepath.Join(dir_abs, new_name))
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
