package def

type Status struct {
	Name                string `json:"name"`
	FullText            string `json:"full_text"`
	Color               string `json:"color"`
	Background          string `json:"background"`
	Separator           bool   `json:"separator"`
	SeparatorBlockWidth int    `json:"separator_block_width"`
	Border              string `json:"border"`
	BorderLeft          int    `json:"border_left"`
	BorderRight         int    `json:"border_right"`
	Urgent              bool   `json:"urgent"`
}

func DefaultStatus() Status {
	return Status{
		Name:                "",
		FullText:            "",
		Color:               "#21222c",
		Background:          "#dcdfe4",
		Separator:           false,
		SeparatorBlockWidth: 0,
		Border:              "#dcdfe4",
		BorderLeft:          5,
		BorderRight:         5,
		Urgent:              false,
	}
}

func (s *Status) Invert() {
  s.Background = "#44475a"
  s.Border = "#44475a"
  s.Color = "#dcdfe4"
}
