package main
import (
    "fmt"
    glru "github.com/hashicorp/golang-lru"
)
func main() {
  size := 4
  recentRatio := 0.5
  ghostRatio := 0.25

  recentSize := int(float64(size) * recentRatio)
	evictSize := int(float64(size) * ghostRatio)

  fmt.Println("recentSize >>", recentSize)
  fmt.Println("evictSize >>", evictSize)

  c, err := glru.New2QParams(size, recentRatio, ghostRatio)
  if err != nil {
    fmt.Print(err)
  }
  // 0 - 0 - 0
  c.Add(1, 1) 
  // 0 - 1 - 0
  if v, ok := c.Get(1); ok {
    fmt.Printf(">> %v\n", v)
  }
  // 1 - 0 - 0

  c.Add(2, 2)
  // 1 - 2 - 0

  c.Add(3, 3)
  // 1 - 3 2 - 0
  fmt.Printf("Lens: %v\n", c.Len())
  if v, ok := c.Get(2); ok {
	  fmt.Printf(">> %v\n", v)
  }

  c.Add(4, 4)
  // 2 1 - 4 3 - 0
  if v, ok := c.Get(4); ok {
	  fmt.Printf(">> %v\n", v)
  }

  c.Add(5, 5)
  // 4 2 - 5 3 - 0
  fmt.Printf("Lens: %v\n", c.Len())

  c.Add(1, 1)
  // 4 2 - 1 5 - 3
  fmt.Printf("Lens: %v\n", c.Len())
  if v, ok := c.Get(3); ok {
	  fmt.Printf(">> %v\n", v)
  }
}
