package Types

type User struct {
	Uid        string `json:"uid"`
	Username   string `json:"username"`
	Nickname   string `json:"nickname"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Admin      bool   `json:"admin"`
	Ban        bool   `json:"ban"`
	Createtime string `json:"createtime"`
}

type Article struct {
	Id         string `json:"id"`
	Author     string `json:"author"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Public     bool   `json:"public"`
	Createtime string `json:"createtime"`
}
