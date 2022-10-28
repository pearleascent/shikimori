package shikimori

import "time"

type AnimeInformationMinimal struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Russian string `json:"russian"`
	Image   struct {
		Original string `json:"original"`
		Preview  string `json:"preview"`
		X96      string `json:"x96"`
		X48      string `json:"x48"`
	} `json:"image"`
	Url           string `json:"url"`
	Kind          string `json:"kind"`
	Score         string `json:"score"`
	Status        string `json:"status"`
	Episodes      int    `json:"episodes"`
	EpisodesAired int    `json:"episodes_aired"`
	AiredOn       string `json:"aired_on"`
	ReleasedOn    string `json:"released_on"`
}

type AnimeInformation struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Russian string `json:"russian"`
	Image   struct {
		Original string `json:"original"`
		Preview  string `json:"preview"`
		X96      string `json:"x96"`
		X48      string `json:"x48"`
	} `json:"image"`
	Url               string   `json:"url"`
	Kind              string   `json:"kind"`
	Score             string   `json:"score"`
	Status            string   `json:"status"`
	Episodes          int      `json:"episodes"`
	EpisodesAired     int      `json:"episodes_aired"`
	AiredOn           string   `json:"aired_on"`
	ReleasedOn        string   `json:"released_on"`
	Rating            string   `json:"rating"`
	English           []string `json:"english"`
	Japanese          []string `json:"japanese"`
	Synonyms          []string `json:"synonyms"`
	LicenseNameRu     string   `json:"license_name_ru"`
	Duration          int      `json:"duration"`
	Description       string   `json:"description"`
	DescriptionHtml   string   `json:"description_html"`
	DescriptionSource string   `json:"description_source"`
	Franchise         string   `json:"franchise"`
	Favoured          bool     `json:"favoured"`
	Anons             bool     `json:"anons"`
	Ongoing           bool     `json:"ongoing"`
	ThreadId          int      `json:"thread_id"`
	TopicId           int      `json:"topic_id"`
	MyanimelistId     int      `json:"myanimelist_id"`
	RatesScoresStats  []struct {
		Name  int `json:"name"`
		Value int `json:"value"`
	} `json:"rates_scores_stats"`
	RatesStatusesStats []struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	} `json:"rates_statuses_stats"`
	UpdatedAt     time.Time   `json:"updated_at"`
	NextEpisodeAt interface{} `json:"next_episode_at"`
	Fansubbers    []string    `json:"fansubbers"`
	Fandubbers    []string    `json:"fandubbers"`
	Licensors     []string    `json:"licensors"`
	Genres        []struct {
		Id      int    `json:"id"`
		Name    string `json:"name"`
		Russian string `json:"russian"`
		Kind    string `json:"kind"`
	} `json:"genres"`
	Studios []struct {
		Id           int    `json:"id"`
		Name         string `json:"name"`
		FilteredName string `json:"filtered_name"`
		Real         bool   `json:"real"`
		Image        string `json:"image"`
	} `json:"studios"`
	Videos []struct {
		Id        int    `json:"id"`
		Url       string `json:"url"`
		ImageUrl  string `json:"image_url"`
		PlayerUrl string `json:"player_url"`
		Name      string `json:"name"`
		Kind      string `json:"kind"`
		Hosting   string `json:"hosting"`
	} `json:"videos"`
	Screenshots []struct {
		Original string `json:"original"`
		Preview  string `json:"preview"`
	} `json:"screenshots"`
}

type AnimeRoles struct {
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
		Url string `json:"url"`
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
		Url string `json:"url"`
	} `json:"person"`
}

type AnimeRelated struct {
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
		Url           string `json:"url"`
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
		Url        string `json:"url"`
		Kind       string `json:"kind"`
		Score      string `json:"score"`
		Status     string `json:"status"`
		Volumes    int    `json:"volumes"`
		Chapters   int    `json:"chapters"`
		AiredOn    string `json:"aired_on"`
		ReleasedOn string `json:"released_on"`
	} `json:"manga"`
}

type Screenshots struct {
	Original string `json:"original"`
	Preview  string `json:"preview"`
}

type Links struct {
	Id         int       `json:"id"`
	Kind       string    `json:"kind"`
	Url        string    `json:"url"`
	Source     string    `json:"source"`
	EntryId    int       `json:"entry_id"`
	EntryType  string    `json:"entry_type"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	ImportedAt time.Time `json:"imported_at"`
}

type AnimeFranchise struct {
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
