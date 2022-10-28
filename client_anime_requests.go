package shikimori

import "encoding/json"

func (c *Client) GetAnimeInformation(id int) (AnimeInformation, error) {
	res, resErr := c.makeRequest(getAnimeByID(id))
	if resErr != nil {
		return AnimeInformation{}, resErr
	}

	animeInfo := AnimeInformation{}
	jsonErr := json.Unmarshal(res, &animeInfo)
	if jsonErr != nil {
		return AnimeInformation{}, resErr
	}

	return animeInfo, nil
}

func (c *Client) GetAnimeRoles(id int) ([]AnimeRoles, error) {
	res, resErr := c.makeRequest(getAnimeRoles(id))
	if resErr != nil {
		return []AnimeRoles{}, resErr
	}

	var animeRoles []AnimeRoles
	jsonErr := json.Unmarshal(res, &animeRoles)
	if jsonErr != nil {
		return []AnimeRoles{}, resErr
	}

	return animeRoles, nil
}

func (c *Client) GetAnimeSimilar(id int) ([]AnimeInformationMinimal, error) {
	res, resErr := c.makeRequest(getAnimeSimilar(id))
	if resErr != nil {
		return []AnimeInformationMinimal{}, resErr
	}

	var animeSimilar []AnimeInformationMinimal
	jsonErr := json.Unmarshal(res, &animeSimilar)
	if jsonErr != nil {
		return []AnimeInformationMinimal{}, resErr
	}

	return animeSimilar, nil
}

func (c *Client) GetAnimeRelated(id int) ([]AnimeRelated, error) {
	res, resErr := c.makeRequest(getAnimeRelated(id))
	if resErr != nil {
		return []AnimeRelated{}, resErr
	}

	var animeRelated []AnimeRelated
	jsonErr := json.Unmarshal(res, &animeRelated)
	if jsonErr != nil {
		return []AnimeRelated{}, resErr
	}

	return animeRelated, nil
}

func (c *Client) GetAnimeScreenshots(id int) ([]Screenshots, error) {
	res, resErr := c.makeRequest(getAnimeScreenshots(id))
	if resErr != nil {
		return []Screenshots{}, resErr
	}

	var animeScreenshots []Screenshots
	jsonErr := json.Unmarshal(res, &animeScreenshots)
	if jsonErr != nil {
		return []Screenshots{}, resErr
	}

	return animeScreenshots, nil
}

func (c *Client) GetAnimeExternalLinks(id int) ([]Links, error) {
	res, resErr := c.makeRequest(getAnimeExternalLinks(id))
	if resErr != nil {
		return []Links{}, resErr
	}

	var animeLinks []Links
	jsonErr := json.Unmarshal(res, &animeLinks)
	if jsonErr != nil {
		return []Links{}, resErr
	}

	return animeLinks, nil
}

func (c *Client) GetAnimeFranchise(id int) (AnimeFranchise, error) {
	res, resErr := c.makeRequest(getAnimeFranchise(id))
	if resErr != nil {
		return AnimeFranchise{}, resErr
	}

	var animeFranchise AnimeFranchise
	jsonErr := json.Unmarshal(res, &animeFranchise)
	if jsonErr != nil {
		return AnimeFranchise{}, resErr
	}

	return animeFranchise, nil
}
