package loader

import "io/ioutil"

func Load(path string) *[]byte {
	file, err := ioutil.ReadFile(path)
	must(err)
	return &file
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
