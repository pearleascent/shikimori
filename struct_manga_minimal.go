package shikimori

type MangaInformationMinimal struct {
	Id      int    `json:"id" bson:"id"`
	Name    string `json:"name" bson:"name"`
	Russian string `json:"russian" bson:"russian"`
	Image   struct {
		Original string `json:"original" bson:"original"`
		Preview  string `json:"preview" bson:"preview"`
		X96      string `json:"x96" bson:"x96"`
		X48      string `json:"x48" bson:"x48"`
	} `json:"image" bson:"image"`
	Url        string `json:"url" bson:"url"`
	Kind       string `json:"kind" bson:"kind"`
	Score      string `json:"score" bson:"score"`
	Status     string `json:"status" bson:"status"`
	Volumes    int    `json:"volumes" bson:"volumes"`
	Chapters   int    `json:"chapters" bson:"chapters"`
	AiredOn    string `json:"aired_on" bson:"aired_on"`
	ReleasedOn string `json:"released_on" bson:"released_on"`
}
