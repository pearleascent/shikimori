package tests

import (
	"github.com/moneydisappointed/shikimori"
	"net/http"
	"testing"
	"time"
)

func TestRanobeInformation(t *testing.T) {
	time.Sleep(500 * time.Millisecond)

	c := shikimori.Init(&http.Client{})
	info, err := c.GetRanobeInformation(123992)

	if err != nil {
		t.Error("Request error:", err)
	}

	if info.Name != "Youkoso Jitsuryoku Shijou Shugi no Kyoushitsu e: 2-nensei-hen" {
		t.Error("Wrong data response data passed")
	}
}

func TestRanobeRoles(t *testing.T) {
	time.Sleep(500 * time.Millisecond)

	c := shikimori.Init(&http.Client{})
	info, err := c.GetRanobeRoles(123992)

	if err != nil {
		t.Error("Request error:", err)
	}

	if info[0].Character.Name != "Kiyotaka Ayanokouji" {
		t.Error("Wrong data response data passed")
	}
}

func TestRanobeSimilar(t *testing.T) {
	time.Sleep(500 * time.Millisecond)

	c := shikimori.Init(&http.Client{})
	info, err := c.GetRanobeSimilar(123992)

	if err != nil {
		t.Error("Request error:", err)
	}

	if info[0].Name != "Kage no Jitsuryokusha ni Naritakute!" {
		t.Error("Wrong data response data passed")
	}
}

func TestRanobeRelated(t *testing.T) {
	time.Sleep(500 * time.Millisecond)

	c := shikimori.Init(&http.Client{})
	info, err := c.GetRanobeRelated(123992)

	if err != nil {
		t.Error("Request error:", err)
	}

	if info[0].Relation != "Prequel" {
		t.Error("Wrong data response data passed")
	}
}

func TestRanobeLinks(t *testing.T) {
	time.Sleep(500 * time.Millisecond)

	c := shikimori.Init(&http.Client{})
	info, err := c.GetRanobeExternalLinks(123992)

	if err != nil {
		t.Error("Request error:", err)
	}

	if info[0].Kind != "official_site" {
		t.Error("Wrong data response data passed")
	}
}

func TestRanobeFranchise(t *testing.T) {
	time.Sleep(500 * time.Millisecond)

	c := shikimori.Init(&http.Client{})
	info, err := c.GetRanobeFranchise(123992)

	if err != nil {
		t.Error("Request error:", err)
	}

	if info.Links[0].TargetId != 89357 {
		t.Error("Wrong data response data passed")
	}
}
