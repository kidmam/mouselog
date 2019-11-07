package trace

type Events struct {
	Url    string  `json:"url"`
	Width  int     `json:"width"`
	Height int     `json:"height"`
	Data   []Event `json:"events"`

	Degrees []float64 `json:"-"`
}

type Event struct {
	Timestamp float64 `json:"timestamp"`
	X         int     `json:"x"`
	Y         int     `json:"y"`
}