package main

import (
	"io/ioutil"
)

func fileDemo() {
	ioutil.WriteFile("./xx.gog", []byte("i am ok\n"), 0644)
}

func main() {
	// fp, err := os.OpenFile("./xx.go", os.O_RDWR|os.O_CREATE, 0644)
	// if err != nil {
	// 	fmt.Printf("open file error:%s", err.Error())
	// }
	// defer fp.Close()
	// fp.Write([]byte{'n', 'b', 'l'})
	fileDemo()
}
