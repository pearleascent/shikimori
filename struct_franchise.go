package shikimori

type Franchise struct {
	Links []struct {
		Id       int    `json:"id"`
		SourceId int    `json:"source_id"`
		TargetId int    `json:"target_id"`
		Source   int    `json:"source"`
		Target   int    `json:"target"`
		Weight   int    `json:"weight"`
		Relation string `json:"relation"`
	} `json:"links"`
	Nodes []struct {
		Id       int    `json:"id"`
		Date     int    `json:"date"`
		Name     string `json:"name"`
		ImageUrl string `json:"image_url"`
		Url      string `json:"url"`
		Year     int    `json:"year"`
		Kind     string `json:"kind"`
		Weight   int    `json:"weight"`
	} `json:"nodes"`
	CurrentId int `json:"current_id"`
}
