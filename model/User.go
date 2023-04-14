package model

type User struct {
	Results []ResultsInfo
}

type ResultsInfo struct {
	Gender string `json:"gender"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
}
