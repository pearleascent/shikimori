package shikimori

import "strconv"

func getAnimeByID(id int) string {
	return "https://shikimori.one/api/animes/" + strconv.Itoa(id)
}

func getAnimeRoles(id int) string {
	return "https://shikimori.one/api/animes/" + strconv.Itoa(id) + "/roles"
}

func getAnimeSimilar(id int) string {
	return "https://shikimori.one/api/animes/" + strconv.Itoa(id) + "/similar"
}

func getAnimeRelated(id int) string {
	return "https://shikimori.one/api/animes/" + strconv.Itoa(id) + "/related"
}

func getAnimeScreenshots(id int) string {
	return "https://shikimori.one/api/animes/" + strconv.Itoa(id) + "/screenshots"
}

func getAnimeFranchise(id int) string {
	return "https://shikimori.one/api/animes/" + strconv.Itoa(id) + "/franchise"
}

func getAnimeExternalLinks(id int) string {
	return "https://shikimori.one/api/animes/" + strconv.Itoa(id) + "/external_links"
}
