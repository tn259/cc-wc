package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

var bytes bool
var lines bool
var words bool
var chars bool
func initArgs() {
  flag.BoolVar(&bytes, "c", false, "Print the number of bytes")
  flag.BoolVar(&lines, "l", false, "Print the number of newlines")
  flag.BoolVar(&words, "w", false, "Print the number of words")
  flag.BoolVar(&chars, "m", false, "Print the number of characters")
  flag.Parse()
}
func handleDefaultOption() {
  if !bytes && !lines && !words && !chars {
    bytes = true
    lines = true
    words = true
    chars = true
  }
}

type counter struct {
  byteCount int
  lineCount int
  wordCount int
  charCount int
}
func NewCounter() *counter {
  return &counter{
    byteCount: 0,
    lineCount: 0,
    wordCount: 0,
    charCount: 0,
  }
}
func (c *counter) incrementByLine(buf []byte) {
  c.byteCount += len(buf)
  c.lineCount++
  c.wordCount += len(strings.Fields(string(buf)))
  c.charCount += utf8.RuneCount(buf)
}
func (c *counter) print() {
  if bytes {
    fmt.Printf("%d ", c.byteCount)
  }
  if lines {
    fmt.Printf("%d ", c.lineCount)
  }
  if words {
    fmt.Printf("%d ", c.wordCount)
  }
  if chars {
    fmt.Printf("%d ", c.charCount)
  }
}

func printCounts(file *os.File) error {  
  reader := bufio.NewReader(file)
  counter := NewCounter()
  for {
    buf, err := reader.ReadBytes('\n')
    if err == io.EOF {
      if len(buf) > 0 {
        counter.incrementByLine(buf)
      }
      break
    }
    if err != nil {
      return err
    }
    counter.incrementByLine(buf)
  }
  counter.print()
  return nil
}

func hasFileArg() bool {
  return flag.NArg() != 1
}
func getFile() (*os.File, error) {
  if hasFileArg() {
    return os.Stdin, nil
  }
  return os.Open(flag.Arg(0))
}

func main() {
	initArgs()
  handleDefaultOption() 

  file, err := getFile()
  if err != nil {
      fmt.Println("Error: Could not open the file:", err)
      os.Exit(1)
  }
  defer file.Close()
  
  if err := printCounts(file); err != nil {
    fmt.Println("Error: Could not print the counts:", err)
    os.Exit(1)
  }

  if hasFileArg() {
    fmt.Println(flag.Arg(0))
  }
}
