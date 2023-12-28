package schema

// Individual link
type Result struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Link    string `json:"link"`
	Details string `json:"details"`
	Types   string `json:"types"`
	Tags    string `json:"tags"`
}

//Search results of all links

type Results struct {
	Results []Result `json:"results"`
}
