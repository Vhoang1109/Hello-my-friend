package response

type Output struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Noti struct {
	Message string `json:"message"`
}
