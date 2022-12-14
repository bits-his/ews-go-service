package app

type Alert struct {
	Headline  string     `json:"headline"`
	Media     []Media    `json:"media"`
	Body      string     `json:"body"`
	Platforms []Platform `json:"platforms"`
	Urls      []Url      `json:"urls"`
	Mails     []Mail     `json:"mails"`
	Phones    []Phone    `json:"phones"`
}

type Url struct {
	Uri string `json:"uri"`
}

type Mail struct {
	Address string `json:"email"`
}

type Media struct {
	Description string `json:"description"`
	Url         string `json:"url"`
}

type Platform struct {
	Name string `json:"name"`
}

type Phone struct {
	Number string `json:"number"`
}
