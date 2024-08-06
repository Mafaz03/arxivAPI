package arxivapi

// import "encoding/xml"

type Feed struct {
	ID      string `bson:"id" xml:"id"`
	Updated string `bson:"updated" xml:"updated"`
	Entry   []struct {
		ID        string `bson:"id" xml:"id"`
		Updated   string `bson:"updated" xml:"updated"`
		Published string `bson:"published" xml:"published"`
		Title     string `bson:"title" xml:"title"`
		NewsTitle string `bson:"newstitle" xml:"newstitle"`
		Image     string `bson:"image" xml:"image"`
		Summary   string `bson:"summary" xml:"summary"`
		Author    []struct {
			Name string `bson:"name" xml:"name"`
		} `bson:"author" xml:"author"`
		Comment []struct {
			Text string `bson:"text" xml:"text"`
		} `bson:"comment" xml:"comment"`
		Link []struct {
			Href string `bson:"href" xml:"href"`
			Rel  string `bson:"rel" xml:"rel"`
			Type string `bson:"type" xml:"type"`
		} `bson:"link" xml:"link"`
		PrimaryCategory []struct {
			Term   string `bson:"term" xml:"term"`
			Scheme string `bson:"scheme" xml:"scheme"`
		} `bson:"primary_category" xml:"primary_category"`
		Category []struct {
			Term   string `bson:"term" xml:"term"`
			Scheme string `bson:"scheme" xml:"scheme"`
		} `bson:"category" xml:"category"`
	} `bson:"entry" xml:"entry"`
}

type Feedjson struct {
	Entry []struct {
		Updated   string `json:"updated"`
		Published string `json:"published"`
		Title     string `json:"title"`
		NewsTitle string `json:"newstitle"`
		Image     string `json:"image"`
		Summary   string `json:"summary"`
		Author    []struct {
			Name string `json:"name"`
		} `json:"author"`
	} `json:"entry"`
}
