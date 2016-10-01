package followmyfriends

import (
	"os"
	"testing"
)

var testToken string

func init() {
	testToken = os.Getenv("access_token")
	if testToken == "" {
		panic("No access token provided")
	}
}

func TestLoadLastActivity(t *testing.T) {
	if testing.Short() {
		t.Skip("Test makes external calls to strava so skip in short mode")
	}
	testObject := StravaActivityLoader{testToken}

	activityId := testObject.lastActivityIdForAthlete(10175532)

	if activityId <= 0 {
		t.Error("Expected a valid activity ID but got ", activityId)
	}
}

func TestNoActivites(t *testing.T) {
	if testing.Short() {
		t.Skip("Test makes external calls to strava so skip in short mode")
	}
	testObject := StravaActivityLoader{testToken}

	activityId := testObject.lastActivityIdForAthlete(99999999)

	if activityId != -1 {
		t.Error("Expected an invalid activity ID but got ", activityId)
	}
}