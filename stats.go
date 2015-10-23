package main


import (
  "io/ioutil"
  "encoding/json"
  "net/http"
  "github.com/moovweb/gokogiri"
  "github.com/moovweb/gokogiri/xpath"
  "regexp"
)

const SAVE_FILE string = "keyboards-update.json"
const URL_EK string = "https://elitekeyboards.com/products.php?sub=topre_keyboards,rftenkeyless"

func extractStats() (keyboards Keyboards) {
  file, e := ioutil.ReadFile(getExecDir() + SAVE_FILE)
  if e == nil {
    json.Unmarshal(file, &keyboards)
  }
  return keyboards
}

func compareStats(nKeyboards, oKeyboards Keyboards) (arr []Keyboard) {
  for key, _ := range nKeyboards.Keyboards {
    _, exist := oKeyboards.Keyboards[key]
    if !exist || (oKeyboards.Keyboards[key].Available !=  nKeyboards.Keyboards[key].Available) {
      arr = append(arr, nKeyboards.Keyboards[key])
    }
  }
  return
}

func saveStats(keyboards Keyboards) {
  b, _ := json.Marshal(keyboards)
  err := ioutil.WriteFile(getExecDir() + SAVE_FILE, b, 0644)
  if err != nil {
    panic(err)
  }
}

func newStats() (keyboards Keyboards) {
  keyboards.Keyboards = make(map[string]Keyboard)
  // fetch and read a web page
  resp, _ := http.Get(URL_EK)
  page, _ := ioutil.ReadAll(resp.Body)

  // parse the web page
  doc, _ := gokogiri.ParseHtml(page)
  defer doc.Free()

  // Extract Keyboard Product
  xps := xpath.Compile("//table[@class='products']/tr[@class='odd']")
  ss, _ := doc.Root().Search(xps)

  // Procompile Regex
  titleXps      := xpath.Compile(".//td[@class='desc']/a")
  descXps       := xpath.Compile(".//td/span[@class='msize']")
  reg, _        := regexp.Compile(`<[-\./\w\":=\s\;]*>`)
  pictureRe, _  := regexp.Compile(`src="\.(.*jpg)" `)
  dataRe, _   := regexp.Compile(`Model: ([\w]+)[\W]+Stock: ([\w]+)[\W]+Price: \$(.*)`)

  // Iter inside of each Keyboard Product
  for _, s := range ss {
    var kb Keyboard
    // Select Image
    pict := pictureRe.FindStringSubmatch(s.InnerHtml())
    kb.Image = pict[1]

    // Select Name
    ss1, _ := s.Search(titleXps)
    for _, s1 := range ss1 {
      kb.Name = s1.InnerHtml()
    }

    // Select Model, Availability, Price
    ss2, _ := s.Search(descXps)
    for _, s2 := range ss2 {
      str := reg.ReplaceAllString(s2.InnerHtml(), "")
      desc := dataRe.FindStringSubmatch(str)
      kb.Model = desc[1]
      kb.Available = desc[2] == "YES"
      kb.Price = desc[3]
    }

    // Insert into the Keboards Collector
    keyboards.Keyboards[kb.Model] = kb
  }
  return
}
