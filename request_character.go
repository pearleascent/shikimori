package shikimori

import (
	"encoding/json"
)

func (c *Client) GetCharacterInformation(id int) (Character, error) {
	res, resErr := c.makeRequest(getByID("characters", id))
	if resErr != nil {
		return Character{}, resErr
	}

	animeInfo := Character{}
	jsonErr := json.Unmarshal(res, &animeInfo)
	if jsonErr != nil {
		return Character{}, jsonErr
	}

	return animeInfo, nil
}
