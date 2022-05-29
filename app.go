package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func app(dir string, mode int, expr string, name_template string, dry_run bool) {
	// get abs path
	dir_abs, err := filepath.Abs(dir)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Work dir:", dir_abs)

	// loop through items in a directory
	files, err := ioutil.ReadDir(dir_abs)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		doRun := (file.IsDir() && (mode == 1)) || (!file.IsDir() && (mode == 0))
		if doRun {
			// If a dir and in dir mode or a file and in file mode, then
			var re = regexp.MustCompile(expr)

			old_name := file.Name()
			new_name := re.ReplaceAllString(old_name, name_template)

			is_skip := old_name == new_name
			if is_skip {
				log.Print(old_name, " -> ", "Skip")
			} else {
				log.Print(old_name, " -> ", new_name)
			}

			if !(dry_run) && !is_skip {
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
