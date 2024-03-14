package api

import (
	"encoding/json"
	"net/http"

	"github.com/baileyjm02/muzz-backend-technical-exercise/match"
	"github.com/baileyjm02/muzz-backend-technical-exercise/swipe"
)

// swipeRequest is the request body for the SwipeUser function.
type swipeRequest struct {
	TargetID   int    `json:"target_id"`
	Preference string `json:"preference"`
}

// swipeResponse is the response body for the SwipeUser function.
type swipeResponse struct {
	Matched bool `json:"matched"`
	MatchID int  `json:"match_id,omitempty"`
}

// SwipeUser allows a user to swipe on another user.
func SwipeUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)

	// Get request data
	req := swipeRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		WriteErrorJSON(w, err)
		return
	}

	// Create a swipe record
	_, err = swipe.CreateSwipe(userID, req.TargetID, req.Preference)
	if err != nil {
		WriteErrorJSON(w, err)
		return
	}

	// Check if the swipe resulted in a match
	matched, err := match.CheckMatch(userID, req.TargetID)
	if err != nil {
		WriteErrorJSON(w, err)
		return
	}

	// If no match, return
	if !matched {
		WriteJSON(w, swipeResponse{
			Matched: false,
		})
		return
	}

	// If match, create a match record
	matchRecord, err := match.CreateMatch(userID, req.TargetID)
	if err != nil {
		WriteErrorJSON(w, err)
		return
	}

	// Return match response
	WriteJSON(w, swipeResponse{
		Matched: true,
		MatchID: matchRecord.ID,
	})
}
