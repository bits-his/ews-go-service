package app

type Alert struct {
	Headline  string     `json:"headline"`
	Media     []Media    `json:"media"`
	Body      string     `json:"body"`
	Platforms []Platform `json:"platforms"`
	Urls      []Url      `json:"urls"`
	Mails     []Mail     `json:"mails"`
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

func (alert Alert) Intf(str string) {

}
