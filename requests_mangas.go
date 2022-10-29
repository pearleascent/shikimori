package shikimori

import (
	"encoding/json"
)

func (c *Client) GetMangaList(page, limit int, order string) ([]MangaInformationMinimal, error) {
	res, resErr := c.makeRequest(getArray("mangas", page, limit, order))
	if resErr != nil {
		return []MangaInformationMinimal{}, resErr
	}

	var animeInfo []MangaInformationMinimal
	jsonErr := json.Unmarshal(res, &animeInfo)
	if jsonErr != nil {
		return []MangaInformationMinimal{}, jsonErr
	}

	return animeInfo, nil
}

func (c *Client) GetMangaInformation(id int) (MangaInformation, error) {
	res, resErr := c.makeRequest(getByID("mangas", id))
	if resErr != nil {
		return MangaInformation{}, resErr
	}

	animeInfo := MangaInformation{}
	jsonErr := json.Unmarshal(res, &animeInfo)
	if jsonErr != nil {
		return MangaInformation{}, jsonErr
	}

	return animeInfo, nil
}

func (c *Client) GetMangaRoles(id int) ([]Roles, error) {
	res, resErr := c.makeRequest(getRoles("mangas", id))
	if resErr != nil {
		return []Roles{}, resErr
	}

	var animeRoles []Roles
	jsonErr := json.Unmarshal(res, &animeRoles)
	if jsonErr != nil {
		return []Roles{}, jsonErr
	}

	return animeRoles, nil
}

func (c *Client) GetMangaSimilar(id int) ([]MangaInformationMinimal, error) {
	res, resErr := c.makeRequest(getSimilar("mangas", id))
	if resErr != nil {
		return []MangaInformationMinimal{}, resErr
	}

	var animeSimilar []MangaInformationMinimal
	jsonErr := json.Unmarshal(res, &animeSimilar)
	if jsonErr != nil {
		return []MangaInformationMinimal{}, jsonErr
	}

	return animeSimilar, nil
}

func (c *Client) GetMangaRelated(id int) ([]Related, error) {
	res, resErr := c.makeRequest(getRelated("mangas", id))
	if resErr != nil {
		return []Related{}, resErr
	}

	var animeRelated []Related
	jsonErr := json.Unmarshal(res, &animeRelated)
	if jsonErr != nil {
		return []Related{}, jsonErr
	}

	return animeRelated, nil
}

func (c *Client) GetMangaExternalLinks(id int) ([]Links, error) {
	res, resErr := c.makeRequest(getExternalLinks("mangas", id))
	if resErr != nil {
		return []Links{}, resErr
	}

	var animeLinks []Links
	jsonErr := json.Unmarshal(res, &animeLinks)
	if jsonErr != nil {
		return []Links{}, jsonErr
	}

	return animeLinks, nil
}

func (c *Client) GetMangaFranchise(id int) (Franchise, error) {
	res, resErr := c.makeRequest(getFranchise("mangas", id))
	if resErr != nil {
		return Franchise{}, resErr
	}

	var animeFranchise Franchise
	jsonErr := json.Unmarshal(res, &animeFranchise)
	if jsonErr != nil {
		return Franchise{}, jsonErr
	}

	return animeFranchise, nil
}
