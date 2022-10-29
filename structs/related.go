package structs

type Related struct {
	Relation        string `json:"relation"`
	RelationRussian string `json:"relation_russian"`
	Anime           struct {
		Id      int    `json:"id"`
		Name    string `json:"name"`
		Russian string `json:"russian"`
		Image   struct {
			Original string `json:"original"`
			Preview  string `json:"preview"`
			X96      string `json:"x96"`
			X48      string `json:"x48"`
		} `json:"image"`
		Url           string `json:"urls.go"`
		Kind          string `json:"kind"`
		Score         string `json:"score"`
		Status        string `json:"status"`
		Episodes      int    `json:"episodes"`
		EpisodesAired int    `json:"episodes_aired"`
		AiredOn       string `json:"aired_on"`
		ReleasedOn    string `json:"released_on"`
	} `json:"anime"`
	Manga struct {
		Id      int    `json:"id"`
		Name    string `json:"name"`
		Russian string `json:"russian"`
		Image   struct {
			Original string `json:"original"`
			Preview  string `json:"preview"`
			X96      string `json:"x96"`
			X48      string `json:"x48"`
		} `json:"image"`
		Url        string `json:"urls.go"`
		Kind       string `json:"kind"`
		Score      string `json:"score"`
		Status     string `json:"status"`
		Volumes    int    `json:"volumes"`
		Chapters   int    `json:"chapters"`
		AiredOn    string `json:"aired_on"`
		ReleasedOn string `json:"released_on"`
	} `json:"manga"`
}
