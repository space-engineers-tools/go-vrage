package vrage

import "log"

// APIAdmin provides access to the /v1/admin API routes.
type APIAdmin struct {
	http *HTTPClient
}

// region /v1/admin/promotedPlayers

// PromotePlayer promotes a player with the specified Steam ID.
func (a *APIAdmin) PromotePlayer(steamID uint64) (APIResponseWithoutData, error) {
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
func (a *APIAdmin) DemotePlayer(steamID uint64) (APIResponseWithoutData, error) {
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

// region /v1/admin/bannedPlayers

// APIAdminBannedPlayersData represents the data returned by the GET /v1/admin/bannedPlayers endpoint.
type APIAdminBannedPlayersData struct {
	BannedPlayers []struct {
		SteamID     uint64 `json:"SteamId"`
		DisplayName string `json:"DisplayName"` // DisplayName can be empty
	} `json:"BannedPlayers"`
}

// BannedPlayers fetches a list of banned players.
func (a *APIAdmin) BannedPlayers() (APIResponseWithData[APIAdminBannedPlayersData], error) {
	var responseStruct APIResponseWithData[APIAdminBannedPlayersData]
	httpResponse, err := a.http.GetV1AdminBannedPlayers()
	if err != nil {
		return responseStruct, err
	}
	defer func() {
		if err := httpResponse.Body.Close(); err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}()

	responseStruct, err = parseResponse[APIResponseWithData[APIAdminBannedPlayersData]](httpResponse)
	if err != nil {
		return responseStruct, err
	}

	return responseStruct, nil
}

// BanPlayer bans a player with the specified Steam ID.
func (a *APIAdmin) BanPlayer(steamID uint64) (APIResponseWithoutData, error) {
	var responseStruct APIResponseWithoutData
	httpResponse, err := a.http.PostV1AdminBannedPlayersSteamID(steamID)
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

// UnbanPlayer unbans a player with the specified Steam ID.
func (a *APIAdmin) UnbanPlayer(steamID uint64) (APIResponseWithoutData, error) {
	var responseStruct APIResponseWithoutData
	httpResponse, err := a.http.DeleteV1AdminBannedPlayersSteamID(steamID)
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
