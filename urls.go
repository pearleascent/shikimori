package shikimori

import "strconv"

func getArray(entry string, page, limit int, order string) string {
	return "https://shikimori.one/api/" + entry + "?page=" + strconv.Itoa(page) + "&limit=" + strconv.Itoa(limit) + "&order=" + order
}

func getByID(entry string, id int) string {
	return "https://shikimori.one/api/" + entry + "/" + strconv.Itoa(id)
}

func getRoles(entry string, id int) string {
	return "https://shikimori.one/api/" + entry + "/" + strconv.Itoa(id) + "/roles"
}

func getSimilar(entry string, id int) string {
	return "https://shikimori.one/api/" + entry + "/" + strconv.Itoa(id) + "/similar"
}

func getRelated(entry string, id int) string {
	return "https://shikimori.one/api/" + entry + "/" + strconv.Itoa(id) + "/related"
}

func getScreenshots(entry string, id int) string {
	return "https://shikimori.one/api/" + entry + "/" + strconv.Itoa(id) + "/screenshots"
}

func getFranchise(entry string, id int) string {
	return "https://shikimori.one/api/" + entry + "/" + strconv.Itoa(id) + "/franchise"
}

func getExternalLinks(entry string, id int) string {
	return "https://shikimori.one/api/" + entry + "/" + strconv.Itoa(id) + "/external_links"
}
