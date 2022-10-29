package shikimori

import (
	"encoding/json"
	"github.com/moneydisappointed/shikimori/structs"
)

func (c *Client) GetAnimeInformation(id int) (structs.AnimeInformation, error) {
	res, resErr := c.makeRequest(getByID("animes", id))
	if resErr != nil {
		return structs.AnimeInformation{}, resErr
	}

	animeInfo := structs.AnimeInformation{}
	jsonErr := json.Unmarshal(res, &animeInfo)
	if jsonErr != nil {
		return structs.AnimeInformation{}, jsonErr
	}

	return animeInfo, nil
}

func (c *Client) GetAnimeRoles(id int) ([]structs.Roles, error) {
	res, resErr := c.makeRequest(getRoles("animes", id))
	if resErr != nil {
		return []structs.Roles{}, resErr
	}

	var animeRoles []structs.Roles
	jsonErr := json.Unmarshal(res, &animeRoles)
	if jsonErr != nil {
		return []structs.Roles{}, jsonErr
	}

	return animeRoles, nil
}

func (c *Client) GetAnimeSimilar(id int) ([]structs.AnimeInformationMinimal, error) {
	res, resErr := c.makeRequest(getSimilar("animes", id))
	if resErr != nil {
		return []structs.AnimeInformationMinimal{}, resErr
	}

	var animeSimilar []structs.AnimeInformationMinimal
	jsonErr := json.Unmarshal(res, &animeSimilar)
	if jsonErr != nil {
		return []structs.AnimeInformationMinimal{}, jsonErr
	}

	return animeSimilar, nil
}

func (c *Client) GetAnimeRelated(id int) ([]structs.Related, error) {
	res, resErr := c.makeRequest(getRelated("animes", id))
	if resErr != nil {
		return []structs.Related{}, resErr
	}

	var animeRelated []structs.Related
	jsonErr := json.Unmarshal(res, &animeRelated)
	if jsonErr != nil {
		return []structs.Related{}, jsonErr
	}

	return animeRelated, nil
}

func (c *Client) GetAnimeScreenshots(id int) ([]structs.Screenshots, error) {
	res, resErr := c.makeRequest(getScreenshots("animes", id))
	if resErr != nil {
		return []structs.Screenshots{}, resErr
	}

	var animeScreenshots []structs.Screenshots
	jsonErr := json.Unmarshal(res, &animeScreenshots)
	if jsonErr != nil {
		return []structs.Screenshots{}, jsonErr
	}

	return animeScreenshots, nil
}

func (c *Client) GetAnimeExternalLinks(id int) ([]structs.Links, error) {
	res, resErr := c.makeRequest(getExternalLinks("animes", id))
	if resErr != nil {
		return []structs.Links{}, resErr
	}

	var animeLinks []structs.Links
	jsonErr := json.Unmarshal(res, &animeLinks)
	if jsonErr != nil {
		return []structs.Links{}, jsonErr
	}

	return animeLinks, nil
}

func (c *Client) GetAnimeFranchise(id int) (structs.Franchise, error) {
	res, resErr := c.makeRequest(getFranchise("animes", id))
	if resErr != nil {
		return structs.Franchise{}, resErr
	}

	var animeFranchise structs.Franchise
	jsonErr := json.Unmarshal(res, &animeFranchise)
	if jsonErr != nil {
		return structs.Franchise{}, jsonErr
	}

	return animeFranchise, nil
}
