package shikimori

import (
	"encoding/json"
)

func (c *Client) GetRanobeList(page, limit int, order string) ([]RanobeInformationMinimal, error) {
	res, resErr := c.makeRequest(getArray("ranobe", page, limit, order))
	if resErr != nil {
		return []RanobeInformationMinimal{}, resErr
	}

	var animeInfo []RanobeInformationMinimal
	jsonErr := json.Unmarshal(res, &animeInfo)
	if jsonErr != nil {
		return []RanobeInformationMinimal{}, jsonErr
	}

	return animeInfo, nil
}

func (c *Client) GetRanobeInformation(id int) (RabobeInformation, error) {
	res, resErr := c.makeRequest(getByID("ranobe", id))
	if resErr != nil {
		return RabobeInformation{}, resErr
	}

	animeInfo := RabobeInformation{}
	jsonErr := json.Unmarshal(res, &animeInfo)
	if jsonErr != nil {
		return RabobeInformation{}, jsonErr
	}

	return animeInfo, nil
}

func (c *Client) GetRanobeRoles(id int) ([]Roles, error) {
	res, resErr := c.makeRequest(getRoles("ranobe", id))
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

func (c *Client) GetRanobeSimilar(id int) ([]RanobeInformationMinimal, error) {
	res, resErr := c.makeRequest(getSimilar("ranobe", id))
	if resErr != nil {
		return []RanobeInformationMinimal{}, resErr
	}

	var animeSimilar []RanobeInformationMinimal
	jsonErr := json.Unmarshal(res, &animeSimilar)
	if jsonErr != nil {
		return []RanobeInformationMinimal{}, jsonErr
	}

	return animeSimilar, nil
}

func (c *Client) GetRanobeRelated(id int) ([]Related, error) {
	res, resErr := c.makeRequest(getRelated("ranobe", id))
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

func (c *Client) GetRanobeExternalLinks(id int) ([]Links, error) {
	res, resErr := c.makeRequest(getExternalLinks("ranobe", id))
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

func (c *Client) GetRanobeFranchise(id int) (Franchise, error) {
	res, resErr := c.makeRequest(getFranchise("ranobe", id))
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
