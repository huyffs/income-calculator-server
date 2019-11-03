package tax

// Band represents a tax band
type Band struct {
	Rate  float64 `json:"rate"`
	Limit int64   `json:"limit"`
}

// Category represents a tax category
type Category struct {
	Title string `json:"title"`
	Bands []Band `json:"bands"`
}

// Year represents a tax year
type Year struct {
	Title      string     `json:"title"`
	Categories []Category `json:"categories"`
}

// Data represents tax data for a region
type Data struct {
	Title string `json:"title"`
	Years []Year `json:"years"`
}
