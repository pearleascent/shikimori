package structs

type Roles struct {
	Roles        []string `json:"roles"`
	RolesRussian []string `json:"roles_russian"`
	Character    struct {
		Id      int    `json:"id"`
		Name    string `json:"name"`
		Russian string `json:"russian"`
		Image   struct {
			Original string `json:"original"`
			Preview  string `json:"preview"`
			X96      string `json:"x96"`
			X48      string `json:"x48"`
		} `json:"image"`
		Url string `json:"urls.go"`
	} `json:"character"`
	Person struct {
		Id      int    `json:"id"`
		Name    string `json:"name"`
		Russian string `json:"russian"`
		Image   struct {
			Original string `json:"original"`
			Preview  string `json:"preview"`
			X96      string `json:"x96"`
			X48      string `json:"x48"`
		} `json:"image"`
		Url string `json:"urls.go"`
	} `json:"person"`
}
