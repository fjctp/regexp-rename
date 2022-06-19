package app

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

func Setup() (string, []string, []string) {
	tmp_dir, err1 := ioutil.TempDir("", "regexp-rename-test-*")
	Check(err1)
	log.Println("Created temp directory: ", tmp_dir)

	test_dirs := []string{
		path.Join(tmp_dir, "test 01"),
		path.Join(tmp_dir, "test 02"),
	}

	test_files := []string{
		path.Join(tmp_dir, "test 01.mp4"),
		path.Join(tmp_dir, "test 02.mp4"),
	}

	for _, test_dir := range test_dirs {
		err2 := os.Mkdir(test_dir, 0755)
		Check(err2)
	}

	for _, test_file := range test_files {
		_, err3 := os.Create(test_file)
		Check(err3)
	}

	return tmp_dir, test_dirs, test_files
}

func TearDown(tmp_dir string) {
	//os.RemoveAll(tmp_dir)
	log.Println("Removed temp directory: ", tmp_dir)
}
