package data

import (
	"testing"
)

func TestValidate(t *testing.T) {
	testJob := &NewJob{
		Position:     "test position",
		Organization: "test org",
		Url:          "https://test.com/",
		Email:        "test@test.com",
	}

	// test valid url format
	result := testJob.Validate(false)
	if result["url"] == "Must provide a valid Url" {
		t.Error("valid url, should have no error - result was=", result["url"])
	}

	// test valid email format
	result = testJob.Validate(false)
	if result["email"] == "Must provide a valid Email" {
		t.Error("valid email, should have no error - result was=", result["email"])
	}

	// test bad url format
	testJob.Url = "https//test.com/"
	result = testJob.Validate(false)
	if result["url"] != "Must provide a valid Url" {
		t.Error("bad url, should show an error - result was=", result["url"])
	}

	// test bad email format
	testJob.Email = "testtest.com"
	result = testJob.Validate(false)
	if result["email"] != "Must provide a valid Email" {
		t.Error("bad email, should show an error - result was=", result["email"])
	}
}
