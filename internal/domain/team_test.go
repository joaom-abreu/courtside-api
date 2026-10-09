package domain

import "testing"

func TestTeam_FullName(t *testing.T) {
	team := Team{City: "Boston", Name: "Celtics"}

	if got := team.FullName(); got != "Boston Celtics" {
		t.Errorf("expected %q, got %q", "Boston Celtics", got)
	}
}

func TestConference_IsValid(t *testing.T) {
	tests := []struct {
		conference Conference
		want       bool
	}{
		{ConferenceEast, true},
		{ConferenceWest, true},
		{Conference("North"), false},
		{Conference(""), false},
	}

	for _, tt := range tests {
		if got := tt.conference.IsValid(); got != tt.want {
			t.Errorf("Conference(%q).IsValid() = %v, want %v", tt.conference, got, tt.want)
		}
	}
}
