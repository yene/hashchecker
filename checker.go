package main

import (
    "crypto/sha1"
    "fmt"
    "os"
    "path/filepath"
    "io/ioutil"
    "bufio"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func shasum(f string) string {
    dat,_ := ioutil.ReadFile(f)
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
    if (info.Size() > 8000) { // skip big files
        return nil
    }
    if info.Mode().IsRegular() {
        if (foundHash(shasum(path))) {
            fmt.Printf("file \"%s\" has matching hash\n", string(path))
        }
    }
    return nil
}

func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

var hashes = make([]string, 1)

func main() {
    path := "./"
    if (len(os.Args) == 2) {
        path = os.Args[1];
    }

    hashes,_ = readLines("hashes.txt")

    err := filepath.Walk(path, travel)
    check(err)
}
