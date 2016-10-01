package followmyfriends

import (
	"flag"
	"testing"
)

func TestFlagParsesAthleteId(t *testing.T) {
	input := NewFlagInput()
	err := flag.Set("id", "1")
	if (err != nil) {
		print("Have error ", err.Error())
		t.Error(err.Error())
	}

	actualId := input.athleteId()
	if actualId != 1 {
		t.Error("Expected athlete ID of 1 to be returned but got ", actualId)
	}
}