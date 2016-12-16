package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Reps struct {
	Results []struct {
		BioguideID  string      `json:"bioguide_id"`
		Birthday    string      `json:"birthday"`
		Chamber     *string     `json:"chamber"`
		ContactForm string      `json:"contact_form"`
		CrpID       string      `json:"crp_id"`
		District    interface{} `json:"district"`
		FacebookID  string      `json:"facebook_id"`
		Fax         string      `json:"fax"`
		FecIds      []string    `json:"fec_ids"`
		FirstName   string      `json:"first_name"`
		Gender      string      `json:"gender"`
		GovtrackID  string      `json:"govtrack_id"`
		IcpsrID     int         `json:"icpsr_id"`
		InOffice    bool        `json:"in_office"`
		LastName    string      `json:"last_name"`
		LisID       string      `json:"lis_id,omitempty"`
		MiddleName  interface{} `json:"middle_name"`
		NameSuffix  interface{} `json:"name_suffix"`
		Nickname    interface{} `json:"nickname"`
		OcEmail     string      `json:"oc_email"`
		OcdID       string      `json:"ocd_id"`
		Office      string      `json:"office"`
		Party       string      `json:"party"`
		Phone       string      `json:"phone"`
		SenateClass int         `json:"senate_class,omitempty"`
		State       string      `json:"state"`
		StateName   string      `json:"state_name"`
		StateRank   string      `json:"state_rank,omitempty"`
		TermEnd     string      `json:"term_end"`
		TermStart   string      `json:"term_start"`
		ThomasID    string      `json:"thomas_id"`
		Title       string      `json:"title"`
		TwitterID   string      `json:"twitter_id"`
		VotesmartID int         `json:"votesmart_id"`
		Website     string      `json:"website"`
		YoutubeID   string      `json:"youtube_id"`
	} `json:"results"`
	Count int `json:"count"`
	Page  struct {
		Count   int `json:"count"`
		PerPage int `json:"per_page"`
		Page    int `json:"page"`
	} `json:"page"`
}

func formatResults(reps *Reps) {
	for _, rep := range reps.Results {
		if *rep.Chamber == "senate" {
			*rep.Chamber = "Senate"
		}
		if *rep.Chamber == "house" {
			*rep.Chamber = "House"
		}
	}
}

func zipLookup(zip string) Reps {
	safeZip := url.QueryEscape(zip)

	url := fmt.Sprintf("https://congress.api.sunlightfoundation.com/legislators/locate?zip=%s", safeZip)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal("Do: ", err)
	}

	defer response.Body.Close()

	var reps Reps
	if err := json.NewDecoder(response.Body).Decode(&reps); err != nil {
		log.Println(err)
	}
	formatResults(&reps)
	return reps
}
