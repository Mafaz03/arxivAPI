package apis

import "encoding/xml"

type Feed struct {
	XMLName xml.Name `xml:"feed"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Link    struct {
		Text string `xml:",chardata"`
		Href string `xml:"href,attr"`
		Rel  string `xml:"rel,attr"`
		Type string `xml:"type,attr"`
	} `xml:"link"`
	Title struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
	} `xml:"title"`
	ID           string `xml:"id"`
	Updated      string `xml:"updated"`
	TotalResults struct {
		Text       string `xml:",chardata"`
		Opensearch string `xml:"opensearch,attr"`
	} `xml:"totalResults"`
	StartIndex struct {
		Text       string `xml:",chardata"`
		Opensearch string `xml:"opensearch,attr"`
	} `xml:"startIndex"`
	ItemsPerPage struct {
		Text       string `xml:",chardata"`
		Opensearch string `xml:"opensearch,attr"`
	} `xml:"itemsPerPage"`
	Entry struct {
		Text      string `xml:",chardata"`
		ID        string `xml:"id"`
		Updated   string `xml:"updated"`
		Published string `xml:"published"`
		Title     string `xml:"title"`
		Summary   string `xml:"summary"`
		Author    []struct {
			Text string `xml:",chardata"`
			Name string `xml:"name"`
		} `xml:"author"`
		Comment struct {
			Text  string `xml:",chardata"`
			Arxiv string `xml:"arxiv,attr"`
		} `xml:"comment"`
		Link []struct {
			Text  string `xml:",chardata"`
			Href  string `xml:"href,attr"`
			Rel   string `xml:"rel,attr"`
			Type  string `xml:"type,attr"`
			Title string `xml:"title,attr"`
		} `xml:"link"`
		PrimaryCategory struct {
			Text   string `xml:",chardata"`
			Arxiv  string `xml:"arxiv,attr"`
			Term   string `xml:"term,attr"`
			Scheme string `xml:"scheme,attr"`
		} `xml:"primary_category"`
		Category struct {
			Text   string `xml:",chardata"`
			Term   string `xml:"term,attr"`
			Scheme string `xml:"scheme,attr"`
		} `xml:"category"`
	} `xml:"entry"`
} 