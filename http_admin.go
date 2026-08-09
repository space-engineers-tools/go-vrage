package vrage

import "net/http"

// region /v1/admin/promotedPlayers

// PostV1AdminPromotedPlayersSteamID promotes a player with the specified Steam ID and returns the HTTP response.
//
//	POST /v1/admin/promotedPlayers/{steam_id}
func (c *HTTPClient) PostV1AdminPromotedPlayersSteamID(steamID string) (*http.Response, error) {
	return c.DoErr(http.MethodPost, "/v1/admin/promotedPlayers/"+steamID, jsonMap(nil), httpHeaders(nil))
}

// DeleteV1AdminPromotedPlayersSteamID demotes a player with the specified Steam ID and returns the HTTP response.
//
//	DELETE /v1/admin/promotedPlayers/{steam_id}
func (c *HTTPClient) DeleteV1AdminPromotedPlayersSteamID(steamID string) (*http.Response, error) {
	return c.DoErr(http.MethodDelete, "/v1/admin/promotedPlayers/"+steamID, jsonMap(nil), httpHeaders(nil))
}

// region /v1/admin/bannedPlayers
// todo: v1/admin/bannedPlayers endpoints

// region /v1/admin/kickedPlayers
// todo: v1/admin/kickedPlayers endpoints
