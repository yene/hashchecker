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

func foundHash(h string) bool {
    for _, value := range hashes {
        if (h == value) {
            return true;
        }
    }
    return false;
}

var hashes = make([]string, 1)

func main() {

    h := shasum("yawe.priv")
    hashes = append(hashes, h)

    fmt.Printf(h)

    visit := func(path string, info os.FileInfo, err error) error {
        if info.IsDir() {
            fmt.Println("dir:  ", path)
        } else {
            fmt.Println("file: ", path)
            if (foundHash(shasum(path))) {
                fmt.Println("found file")
            }
        }
        return nil
    }

    err := filepath.Walk("./", visit)
    check(err)
}
