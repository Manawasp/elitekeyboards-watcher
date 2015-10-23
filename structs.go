package main

type Keyboards struct {
  Keyboards map[string]Keyboard
}

type Keyboard struct {
  Name string
  Price string
  Model string
  Image string
  Available bool
}
