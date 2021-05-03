package repository

type Prop struct {
	Id    int    `json:"id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Widget struct {
	Name  string `json:"name"`
	Props []Prop `json:"props"`
}

type Screen struct {
	Id      int      `json:"id"`
	Type    string   `json:"type"`
	Props   []Prop   `json:"props"`
	Widgets []Widget `json:"widgets"`
}

type App struct {
	Props   []Prop   `json:"props"`
	Screens []Screen `json:"screens"`
}

type Project struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	App  App    `json:"app"`
}

type UserData struct {
	Id       string    `json:"id"`
	Email    string    `json:"email"`
	Projects []Project `json:"projects"`
}
