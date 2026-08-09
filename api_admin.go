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

// APIAdminBannedPlayer represents a banned player.
type APIAdminBannedPlayer struct {
	SteamID     uint64 `json:"SteamId"`
	DisplayName string `json:"DisplayName"` // DisplayName can be empty
}

// APIAdminBannedPlayersData represents the data returned by the GET /v1/admin/bannedPlayers endpoint.
type APIAdminBannedPlayersData struct {
	BannedPlayers []APIAdminBannedPlayer `json:"BannedPlayers"`
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

// region /v1/admin/kickedPlayers

// APIAdminKickedPlayer represents a kicked player.
type APIAdminKickedPlayer struct {
	SteamID uint64 `json:"SteamID"`
	// DisplayName is the display name of the player. It can be an empty string.
	DisplayName string `json:"DisplayName"`
	// Time is the remaining kick duration in milliseconds.
	// It gets negative when the player is allowed to join again and the entry is deleted when the player joins again.
	Time int64 `json:"Time"`
}

// CanJoin returns true if the player is allowed to join again (Time <= 0).
func (p APIAdminKickedPlayer) CanJoin() bool {
	return p.Time <= 0
}

// APIAdminKickedPlayersData represents the data returned by the GET /v1/admin/kickedPlayers endpoint.
type APIAdminKickedPlayersData struct {
	KickedPlayers []APIAdminKickedPlayer `json:"KickedPlayers"`
}

// KickedPlayers fetches a list of kicked players.
//
//	GET /v1/admin/kickedPlayers
func (a *APIAdmin) KickedPlayers() (APIResponseWithData[APIAdminKickedPlayersData], error) {
	var responseStruct APIResponseWithData[APIAdminKickedPlayersData]
	httpResponse, err := a.http.GetV1AdminKickedPlayers()
	if err != nil {
		return responseStruct, err
	}
	defer func() {
		if err := httpResponse.Body.Close(); err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}()

	responseStruct, err = parseResponse[APIResponseWithData[APIAdminKickedPlayersData]](httpResponse)
	if err != nil {
		return responseStruct, err
	}

	return responseStruct, nil
}

// KickPlayer kicks a player with the specified Steam ID.
//
// This will prevent the player from joining the server for 5 minutes.
//
//	POST /v1/admin/kickedPlayers/{steam_id}
func (a *APIAdmin) KickPlayer(steamID uint64) (APIResponseWithoutData, error) {
	var responseStruct APIResponseWithoutData
	httpResponse, err := a.http.PostV1AdminKickedPlayersSteamID(steamID)
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

// UnkickPlayer un-kicks a player with the specified Steam ID.
//
//	DELETE /v1/admin/kickedPlayers/{steam_id}
func (a *APIAdmin) UnkickPlayer(steamID uint64) (APIResponseWithoutData, error) {
	var responseStruct APIResponseWithoutData
	httpResponse, err := a.http.DeleteV1AdminKickedPlayersSteamID(steamID)
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
