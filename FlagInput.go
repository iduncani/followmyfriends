package followmyfriends

import "flag"

type FlagInput struct {
	id int64
}

func NewFlagInput() *FlagInput {
	flagInput := FlagInput{}
	flag.Int64Var(&flagInput.id, "id", 0, "Athlete ID")
	return &flagInput
}

func (input *FlagInput) athleteId() int64 {
	if !flag.Parsed() {
		flag.Parse()
	}
	return input.id
}
