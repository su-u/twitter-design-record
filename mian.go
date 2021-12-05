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
    "テキスト+Twitter Large Card": "https://twitter.com/TwitterDevJP/status/1354226852466835459",
    "テキスト+Twitter Large Card（description）": "https://twitter.com/TwitterDevJP/status/971972766852251648",
    "テキスト+Twitter Small Card": "https://twitter.com/TwitterDevJP/status/941468621880008704",
    "テキスト+Twitter Small Card（プレビュー画像、description無し）": "https://twitter.com/TwitterDevJP/status/1300589123883773952",
    "引用ツイート+画像": "https://twitter.com/TwitterDevJP/status/1172324315557023745",
    "4毎の画像": "https://twitter.com/contain1101/status/1428169069363175430",
    "テキスト+URL+画像": "https://twitter.com/TwitterDevJP/status/732819781217325057",
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