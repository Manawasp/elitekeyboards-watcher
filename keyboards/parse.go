package keyboards

import (
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xpath"
)

func WebParse(url string) (keyboards Keyboards) {
	keyboards = make(map[string]Keyboard)
	// fetch and read a web page
	resp, _ := http.Get(url)
	page, _ := ioutil.ReadAll(resp.Body)

	// parse the web page
	doc, _ := gokogiri.ParseHtml(page)
	defer doc.Free()

	// Extract Keyboard Product
	xps := xpath.Compile("//table[@class='products']/tr[@class='odd']")
	ss, _ := doc.Root().Search(xps)

	// Precompile Regex
	titleXps := xpath.Compile(".//td[@class='desc']/a")
	descXps := xpath.Compile(".//td/span[@class='msize']")
	reg, _ := regexp.Compile(`<[-\./\w\":=\s\;]*>`)
	pictureRe, _ := regexp.Compile(`src="\.(.*jpg)" `)
	dataRe, _ := regexp.Compile(`Model: ([\w]+)[\W]+Stock: ([\w]+)[\W]+Price: \$(.*)`)

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
		keyboards[kb.Model] = kb
	}
	return
}
