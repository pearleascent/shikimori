package shikimori

import "time"

type AnimeInformation struct {
	Id      int    `json:"id"  bson:"id"`
	Name    string `json:"name"  bson:"name"`
	Russian string `json:"russian"  bson:"russian"`
	Image   struct {
		Original string `json:"original"  bson:"original"`
		Preview  string `json:"preview"  bson:"preview"`
		X96      string `json:"x96"  bson:"x96"`
		X48      string `json:"x48"  bson:"x48"`
	} `json:"image"  bson:"image"`
	Url               string   `json:"url"  bson:"url"`
	Kind              string   `json:"kind"  bson:"kind"`
	Score             string   `json:"score"  bson:"score"`
	Status            string   `json:"status"  bson:"status"`
	Episodes          int      `json:"episodes"  bson:"episodes"`
	EpisodesAired     int      `json:"episodes_aired"  bson:"episodes_aired"`
	AiredOn           string   `json:"aired_on"  bson:"aired_on"`
	ReleasedOn        string   `json:"released_on"  bson:"released_on"`
	Rating            string   `json:"rating"  bson:"rating"`
	English           []string `json:"english"  bson:"english"`
	Japanese          []string `json:"japanese"  bson:"japanese"`
	Synonyms          []string `json:"synonyms"  bson:"synonyms"`
	LicenseNameRu     string   `json:"license_name_ru"  bson:"license_name_ru"`
	Duration          int      `json:"duration" bson:"duration"`
	Description       string   `json:"description"  bson:"description"`
	DescriptionHtml   string   `json:"description_html"  bson:"description_html"`
	DescriptionSource string   `json:"description_source"  bson:"description_source"`
	Franchise         string   `json:"franchise"  bson:"franchise"`
	Favoured          bool     `json:"favoured"  bson:"favoured"`
	Anons             bool     `json:"anons"  bson:"anons"`
	Ongoing           bool     `json:"ongoing"  bson:"ongoing"`
	ThreadId          int      `json:"thread_id"  bson:"thread_id"`
	TopicId           int      `json:"topic_id" bson:"topic_id"`
	MyanimelistId     int      `json:"myanimelist_id"  bson:"myanimelist_id"`
	RatesScoresStats  []struct {
		Name  int `json:"name"  bson:"name"`
		Value int `json:"value"  bson:"value"`
	} `json:"rates_scores_stats"  bson:"rates_scores_stats"`
	RatesStatusesStats []struct {
		Name  string `json:"name"  bson:"name"`
		Value int    `json:"value"  bson:"value"`
	} `json:"rates_statuses_stats" bson:"rates_statuses_stats"`
	UpdatedAt     time.Time   `json:"updated_at" bson:"updated_at"`
	NextEpisodeAt interface{} `json:"next_episode_at" bson:"next_episode_at"`
	Fansubbers    []string    `json:"fansubbers" bson:"fansubbers"`
	Fandubbers    []string    `json:"fandubbers" bson:"fandubbers"`
	Licensors     []string    `json:"licensors" bson:"licensors"`
	Genres        []struct {
		Id      int    `json:"id" bson:"id"`
		Name    string `json:"name" bson:"name"`
		Russian string `json:"russian" bson:"russian"`
		Kind    string `json:"kind" bson:"kind"`
	} `json:"genres" bson:"genres"`
	Studios []struct {
		Id           int    `json:"id" bson:"id"`
		Name         string `json:"name" bson:"name"`
		FilteredName string `json:"filtered_name"  bson:"filtered_name"`
		Real         bool   `json:"real" bson:"real"`
		Image        string `json:"image" bson:"image"`
	} `json:"studios" bson:"studios"`
	Videos []struct {
		Id        int    `json:"id" bson:"id"`
		Url       string `json:"url" bson:"url"`
		ImageUrl  string `json:"image_url" bson:"image_url"`
		PlayerUrl string `json:"player_url" bson:"player_url"`
		Name      string `json:"name" bson:"name"`
		Kind      string `json:"kind" bson:"kind"`
		Hosting   string `json:"hosting" bson:"hosting"`
	} `json:"videos" bson:"videos"`
	Screenshots []struct {
		Original string `json:"original" bson:"original"`
		Preview  string `json:"preview" bson:"preview"`
	} `json:"screenshots" bson:"screenshots"`
}
