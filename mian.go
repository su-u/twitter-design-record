package main

import (
  "github.com/sclevine/agouti"
  "log"
  "os"
  "time"
)

const layout = "2006-01-02 15-04-05"

func main() {
  urls := map[string]string{
    "4毎の画像": "https://twitter.com/contain1101/status/1428169069363175430",
    "FaceBook Appページ": "https://twitter.com/contain1101/status/1460774339511341056",
  }

  t := time.Now().Format(layout)
  if err := os.MkdirAll("twitter/" + t, 0777); err != nil {
    log.Fatal(err)
  }

  // Chromeを利用することを宣言
  d := agouti.ChromeDriver()
  d.Start()
  defer d.Stop()
  p, _ := d.NewPage()

  for k, v := range urls {
    fName := "twitter/" + t + "/" + k + ".png"

    err := p.Navigate(v)
    if err != nil {
      log.Println("Failure: " + fName)
      continue
    }
    time.Sleep(3 * time.Second)

    p.Screenshot(fName)
    log.Println("Done: " + fName)
  }
}