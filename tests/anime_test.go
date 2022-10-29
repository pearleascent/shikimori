package tests

import (
	"github.com/moneydisappointed/shikimori"
	"net/http"
	"testing"
	"time"
)

func TestAnimeList(t *testing.T) {
	time.Sleep(500 * time.Millisecond)

	c := shikimori.Init(&http.Client{})
	info, err := c.GetAnimeList(1, 50, "id")

	if err != nil {
		t.Error("Request error:", err)
	}

	if len(info) != 50 {
		t.Error("Wrong data response data passed")
	}
}

func TestAnimeInformation(t *testing.T) {
	time.Sleep(500 * time.Millisecond)

	c := shikimori.Init(&http.Client{})
	info, err := c.GetAnimeInformation(1)

	if err != nil {
		t.Error("Request error:", err)
	}

	if info.Name != "Cowboy Bebop" {
		t.Error("Wrong data response data passed")
	}
}

func TestAnimeRoles(t *testing.T) {
	time.Sleep(500 * time.Millisecond)

	c := shikimori.Init(&http.Client{})
	info, err := c.GetAnimeRoles(1)

	if err != nil {
		t.Error("Request error:", err)
	}

	if info[0].Character.Name != "Jet Black" {
		t.Error("Wrong data response data passed")
	}
}

func TestAnimeSimilar(t *testing.T) {
	time.Sleep(500 * time.Millisecond)

	c := shikimori.Init(&http.Client{})
	info, err := c.GetAnimeSimilar(1)

	if err != nil {
		t.Error("Request error:", err)
	}

	if info[0].Name != "Samurai Champloo" {
		t.Error("Wrong data response data passed")
	}
}

func TestAnimeRelated(t *testing.T) {
	time.Sleep(500 * time.Millisecond)

	c := shikimori.Init(&http.Client{})
	info, err := c.GetAnimeRelated(1)

	if err != nil {
		t.Error("Request error:", err)
	}

	if info[0].Relation != "Adaptation" {
		t.Error("Wrong data response data passed")
	}
}

func TestAnimeScreenshots(t *testing.T) {
	time.Sleep(500 * time.Millisecond)

	c := shikimori.Init(&http.Client{})
	info, err := c.GetAnimeScreenshots(1)

	if err != nil {
		t.Error("Request error:", err)
	}

	if info[0].Original != "/system/screenshots/original/b1276021bfb0fc514c35bb106361dd2bcb20f967.jpg?1521803518" {
		t.Error("Wrong data response data passed")
	}
}

func TestAnimeLinks(t *testing.T) {
	time.Sleep(500 * time.Millisecond)

	c := shikimori.Init(&http.Client{})
	info, err := c.GetAnimeExternalLinks(1)

	if err != nil {
		t.Error("Request error:", err)
	}

	if info[0].Kind != "official_site" {
		t.Error("Wrong data response data passed")
	}
}

func TestAnimeFranchise(t *testing.T) {
	time.Sleep(500 * time.Millisecond)

	c := shikimori.Init(&http.Client{})
	info, err := c.GetAnimeFranchise(1)

	if err != nil {
		t.Error("Request error:", err)
	}

	if info.Links[0].TargetId != 5 {
		t.Error("Wrong data response data passed")
	}
}
