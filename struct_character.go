package shikimori

import "time"

type Character struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Russian string `json:"russian"`
	Image   struct {
		Original string `json:"original"`
		Preview  string `json:"preview"`
		X96      string `json:"x96"`
		X48      string `json:"x48"`
	} `json:"image"`
	Url               string    `json:"url"`
	Altname           string    `json:"altname"`
	Japanese          string    `json:"japanese"`
	Description       string    `json:"description"`
	DescriptionHtml   string    `json:"description_html"`
	DescriptionSource string    `json:"description_source"`
	Favoured          bool      `json:"favoured"`
	ThreadId          int       `json:"thread_id"`
	TopicId           int       `json:"topic_id"`
	UpdatedAt         time.Time `json:"updated_at"`
	Seyu              []struct {
		Id      int    `json:"id"`
		Name    string `json:"name"`
		Russian string `json:"russian"`
		Image   struct {
			Original string `json:"original"`
			Preview  string `json:"preview"`
			X96      string `json:"x96"`
			X48      string `json:"x48"`
		} `json:"image"`
		Url string `json:"url"`
	} `json:"seyu"`
	Animes []struct {
		Id      int    `json:"id"`
		Name    string `json:"name"`
		Russian string `json:"russian"`
		Image   struct {
			Original string `json:"original"`
			Preview  string `json:"preview"`
			X96      string `json:"x96"`
			X48      string `json:"x48"`
		} `json:"image"`
		Url           string   `json:"url"`
		Kind          string   `json:"kind"`
		Score         string   `json:"score"`
		Status        string   `json:"status"`
		Episodes      int      `json:"episodes"`
		EpisodesAired int      `json:"episodes_aired"`
		AiredOn       string   `json:"aired_on"`
		ReleasedOn    string   `json:"released_on"`
		Roles         []string `json:"roles"`
		Role          string   `json:"role"`
	} `json:"animes"`
	Mangas []struct {
		Id      int    `json:"id"`
		Name    string `json:"name"`
		Russian string `json:"russian"`
		Image   struct {
			Original string `json:"original"`
			Preview  string `json:"preview"`
			X96      string `json:"x96"`
			X48      string `json:"x48"`
		} `json:"image"`
		Url        string   `json:"url"`
		Kind       string   `json:"kind"`
		Score      string   `json:"score"`
		Status     string   `json:"status"`
		Volumes    int      `json:"volumes"`
		Chapters   int      `json:"chapters"`
		AiredOn    string   `json:"aired_on"`
		ReleasedOn string   `json:"released_on"`
		Roles      []string `json:"roles"`
		Role       string   `json:"role"`
	} `json:"mangas"`
}
