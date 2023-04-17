package model

type User struct {
	Results []ResultsInfo
	Info    InfoInfo
}

type ResultsInfo struct {
	Gender     string `json:"gender"`
	Name       NameInfo
	Location   LocationInfo
	Email      string `json:"email"`
	Login      LoginInfo
	Dob        DobInfo
	Registered RegisteredInfo
	Phone      string `json:"phone"`
	Cell       string `json:"cell"`
	Id         IdInfo
	Picture    PictureInfo
	Nat        string `json:"nat"`
}

type NameInfo struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last  string `json:"last"`
}

type LocationInfo struct {
	Street      StreetInfo
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	Postcode    int32  `json:"postcode"`
	Coordinates CoordinatesInfo
	Timezone    TimezoneInfo
}

type StreetInfo struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
}

type CoordinatesInfo struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type TimezoneInfo struct {
	Offset      string `json:"offset"`
	Description string `json:"description"`
}

type LoginInfo struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	Md5      string `json:"md5"`
	Sha1     string `json:"sha1"`
	Sha256   string `json:"sha256"`
}

type DobInfo struct {
	Date string `json:"date"`
	Age  int    `json:"age"`
}

type RegisteredInfo struct {
	Date string `json:"date"`
	Age  int    `json:"age"`
}

type IdInfo struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type PictureInfo struct {
	Large     string `json:"large"`
	Medium    string `json:"medium"`
	Thumbnail string `json:"thumbnail"`
}

type InfoInfo struct {
	Seed    string `json:"seed"`
	Results int    `json:"results"`
	Page    int    `json:"page"`
	Version string `json:"version"`
}
