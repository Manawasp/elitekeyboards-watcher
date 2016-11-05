package keyboards

import (
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/jbowtie/gokogiri"
	"github.com/jbowtie/gokogiri/xpath"
)

// WebParse retrieve html from the given url and look for keyboards
// note: this "parser" only work with the current elitekeyboards website at
// this day 05 Nomvember 2016
func WebParse(url string) (*State, error) {
	state := &State{}
	state.Keyboards = make(map[string]Keyboard)
	// fetch and read a web page
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	page, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// parse the web page
	doc, err := gokogiri.ParseHtml(page)
	if err != nil {
		return nil, err
	}
	defer doc.Free()

	// Extract Keyboard Product
	xps := xpath.Compile("//table[@class='products']/tr[@class='odd']")
	ss, err := doc.Root().Search(xps)
	if err != nil {
		return nil, err
	}

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
		state.Keyboards[kb.Model] = kb
	}
	return state, nil
}
