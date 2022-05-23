package entity

type User struct {
	ID    uint
	Name  string
	Tasks []Task `auto:"true"`
}
