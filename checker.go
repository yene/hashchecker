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

func travel(path string, info os.FileInfo, err error) error {
    if info.Mode().IsRegular() {
        if (foundHash(shasum(path))) {
            fmt.Println("file: ", path)
            fmt.Println("found matching hash")
        }
    }
    return nil
}


var hashes = make([]string, 1)

func main() {

    hashes = append(hashes, "efcc47c3fd5806515a270d6fa0bbe4cc7353eabc")

    err := filepath.Walk("./", travel)
    check(err)
}
