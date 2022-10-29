package tests

import (
	"github.com/moneydisappointed/shikimori"
	"net/http"
	"testing"
	"time"
)

func TestMangaInformation(t *testing.T) {
	time.Sleep(500 * time.Millisecond)

	c := shikimori.Init(&http.Client{})
	info, err := c.GetMangaInformation(2)

	if err != nil {
		t.Error("Request error:", err)
	}

	if info.Name != "Berserk" {
		t.Error("Wrong data response data passed")
	}
}

func TestMangaRoles(t *testing.T) {
	time.Sleep(500 * time.Millisecond)

	c := shikimori.Init(&http.Client{})
	info, err := c.GetMangaRoles(2)

	if err != nil {
		t.Error("Request error:", err)
	}

	if info[0].Character.Name != "Casca" {
		t.Error("Wrong data response data passed")
	}
}

func TestMangaSimilar(t *testing.T) {
	time.Sleep(500 * time.Millisecond)

	c := shikimori.Init(&http.Client{})
	info, err := c.GetMangaSimilar(2)

	if err != nil {
		t.Error("Request error:", err)
	}

	if info[0].Name != "Claymore" {
		t.Error("Wrong data response data passed")
	}
}

func TestMangaRelated(t *testing.T) {
	time.Sleep(500 * time.Millisecond)

	c := shikimori.Init(&http.Client{})
	info, err := c.GetMangaRelated(2)

	if err != nil {
		t.Error("Request error:", err)
	}

	if info[0].Relation != "Other" {
		t.Error("Wrong data response data passed")
	}
}

func TestMangaLinks(t *testing.T) {
	time.Sleep(500 * time.Millisecond)

	c := shikimori.Init(&http.Client{})
	info, err := c.GetMangaExternalLinks(2)

	if err != nil {
		t.Error("Request error:", err)
	}

	if info[0].Kind != "official_site" {
		t.Error("Wrong data response data passed")
	}
}

func TestMangaFranchise(t *testing.T) {
	time.Sleep(500 * time.Millisecond)

	c := shikimori.Init(&http.Client{})
	info, err := c.GetMangaFranchise(2)

	if err != nil {
		t.Error("Request error:", err)
	}

	if info.Links[0].TargetId != 92299 {
		t.Error("Wrong data response data passed")
	}
}
