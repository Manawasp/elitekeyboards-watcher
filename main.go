package main

import (
  "net/http"
  "io/ioutil"
  "fmt"
  "github.com/moovweb/gokogiri"
  "github.com/moovweb/gokogiri/xpath"
  "regexp"
)

func main() {
  m := make(map[string]bool)
  // fetch and read a web page
  resp, _ := http.Get("https://elitekeyboards.com/products.php?sub=topre_keyboards")
  page, _ := ioutil.ReadAll(resp.Body)

  // parse the web page
  doc, _ := gokogiri.ParseHtml(page)
  defer doc.Free()

  xps := xpath.Compile("//table[@class='products']/tr[@class='odd']/td/span[@class='msize']")
  ss, _ := doc.Root().Search(xps)

  // Procompile Regex
  reg, _  := regexp.Compile(`<[-\./\w\":=\s\;]*>`)
  re, _   := regexp.Compile(`Model: ([\w]+)[\W]+Stock: ([\w]+)[\W]+Price: \$(.*)`)

  for _, s := range ss {
    str := reg.ReplaceAllString(s.InnerHtml(), "")

    result := re.FindStringSubmatch(str)
    m[result[1]] = result[2] == "YES"
  }

  if _, key := m["SE19E0"]; key {
    if m["SE19E0"] {
      fmt.Println("available")
    } else {
      fmt.Println("Not available")
    }
    sendEmail()
  }

  // TODO: SAVE information
  // TODO: compare information
  // TODO: send email
  fmt.Println(m)
  // perform operations on the parsed page -- consult the tests for examples
  // fmt.Println(doc.String())
}