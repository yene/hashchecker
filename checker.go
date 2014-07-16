package main

import (
    "crypto/sha1"
    "fmt"
    "os"
    "path/filepath"
    "io/ioutil"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func shasum(f string) string {
    dat, err := ioutil.ReadFile(f)
    check(err)
    hasher := sha1.New()
    hasher.Write(dat)
    return fmt.Sprintf("%x", hasher.Sum(nil))
}


var hashes = make([]string, 1)

func main() {
	//hash := "efcc47c3fd5806515a270d6fa0bbe4cc7353eabc"

    h := shasum("yawe.priv")
    hashes = append(hashes, h)

    fmt.Printf(h)

    visit := func(path string, info os.FileInfo, err error) error {
        if info.IsDir() {
            fmt.Println("dir:  ", path)
        } else {
            fmt.Println("file: ", path)
        }
        return nil
    }

    err := filepath.Walk("./", visit)
    check(err)
}
