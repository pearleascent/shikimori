package shikimori

import (
	"encoding/json"
)

func (c *Client) GetAnimeInformation(id int) (AnimeInformation, error) {
	res, resErr := c.makeRequest(getByID("animes", id))
	if resErr != nil {
		return AnimeInformation{}, resErr
	}

	animeInfo := AnimeInformation{}
	jsonErr := json.Unmarshal(res, &animeInfo)
	if jsonErr != nil {
		return AnimeInformation{}, jsonErr
	}

	return animeInfo, nil
}

func (c *Client) GetAnimeRoles(id int) ([]Roles, error) {
	res, resErr := c.makeRequest(getRoles("animes", id))
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

func (c *Client) GetAnimeSimilar(id int) ([]AnimeInformationMinimal, error) {
	res, resErr := c.makeRequest(getSimilar("animes", id))
	if resErr != nil {
		return []AnimeInformationMinimal{}, resErr
	}

	var animeSimilar []AnimeInformationMinimal
	jsonErr := json.Unmarshal(res, &animeSimilar)
	if jsonErr != nil {
		return []AnimeInformationMinimal{}, jsonErr
	}

	return animeSimilar, nil
}

func (c *Client) GetAnimeRelated(id int) ([]Related, error) {
	res, resErr := c.makeRequest(getRelated("animes", id))
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

func (c *Client) GetAnimeScreenshots(id int) ([]Screenshots, error) {
	res, resErr := c.makeRequest(getScreenshots("animes", id))
	if resErr != nil {
		return []Screenshots{}, resErr
	}

	var animeScreenshots []Screenshots
	jsonErr := json.Unmarshal(res, &animeScreenshots)
	if jsonErr != nil {
		return []Screenshots{}, jsonErr
	}

	return animeScreenshots, nil
}

func (c *Client) GetAnimeExternalLinks(id int) ([]Links, error) {
	res, resErr := c.makeRequest(getExternalLinks("animes", id))
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

func (c *Client) GetAnimeFranchise(id int) (Franchise, error) {
	res, resErr := c.makeRequest(getFranchise("animes", id))
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
