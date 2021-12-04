package main

import (
  "fmt"

  "github.com/sclevine/agouti"
)

func main() {
  // Chromeを利用することを宣言
  d := agouti.ChromeDriver()
  d.Start()
  defer d.Stop()
  p, _ := d.NewPage()

  // 自動操作
  p.Navigate("https://qiita.com/")
  fmt.Println(
    p.Title())
  p.Screenshot("img/Screenshot01.png")

  p.FindByLink("もっと詳しく").Click()
  fmt.Println(p.Title())
  p.Screenshot("img/Screenshot02.png")
}