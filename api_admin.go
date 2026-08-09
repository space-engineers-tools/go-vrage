package vrage

import "log"

// APIAdmin provides access to the /v1/admin API routes.
type APIAdmin struct {
	http *HTTPClient
}

// region /v1/admin/promotedPlayers

// PromotePlayer promotes a player with the specified Steam ID.
func (a *APIAdmin) PromotePlayer(steamID string) (APIResponseWithoutData, error) {
	var responseStruct APIResponseWithoutData
	httpResponse, err := a.http.PostV1AdminPromotedPlayersSteamID(steamID)
	if err != nil {
		return responseStruct, err
	}
	defer func() {
		if err := httpResponse.Body.Close(); err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}()

	responseStruct, err = parseResponse[APIResponseWithoutData](httpResponse)
	if err != nil {
		return responseStruct, err
	}

	return responseStruct, nil
}

// DemotePlayer demotes a player with the specified Steam ID.
func (a *APIAdmin) DemotePlayer(steamID string) (APIResponseWithoutData, error) {
	var responseStruct APIResponseWithoutData
	httpResponse, err := a.http.DeleteV1AdminPromotedPlayersSteamID(steamID)
	if err != nil {
		return responseStruct, err
	}
	defer func() {
		if err := httpResponse.Body.Close(); err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}()

	responseStruct, err = parseResponse[APIResponseWithoutData](httpResponse)
	if err != nil {
		return responseStruct, err
	}

	return responseStruct, nil
}
