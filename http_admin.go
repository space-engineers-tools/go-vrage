package vrage

import (
	"fmt"
	"net/http"
)

// region /v1/admin/promotedPlayers

// PostV1AdminPromotedPlayersSteamID promotes a player with the specified Steam ID and returns the HTTP response.
//
//	POST /v1/admin/promotedPlayers/{steam_id}
func (c *HTTPClient) PostV1AdminPromotedPlayersSteamID(steamID uint64) (*http.Response, error) {
	return c.DoErr(http.MethodPost, fmt.Sprintf("/v1/admin/promotedPlayers/%d", steamID), jsonMap(nil), httpHeaders(nil))
}

// DeleteV1AdminPromotedPlayersSteamID demotes a player with the specified Steam ID and returns the HTTP response.
//
//	DELETE /v1/admin/promotedPlayers/{steam_id}
func (c *HTTPClient) DeleteV1AdminPromotedPlayersSteamID(steamID uint64) (*http.Response, error) {
	return c.DoErr(http.MethodDelete, fmt.Sprintf("/v1/admin/promotedPlayers/%d", steamID), jsonMap(nil), httpHeaders(nil))
}

// region /v1/admin/bannedPlayers

// GetV1AdminBannedPlayers fetches a list of banned players and returns the HTTP response.
//
//	GET /v1/admin/bannedPlayers
func (c *HTTPClient) GetV1AdminBannedPlayers() (*http.Response, error) {
	return c.DoErr(http.MethodGet, "/v1/admin/bannedPlayers", jsonMap(nil), httpHeaders(nil))
}

// PostV1AdminBannedPlayersSteamID bans a player with the specified Steam ID and returns the HTTP response.
//
//	POST /v1/admin/bannedPlayers/{steam_id}
func (c *HTTPClient) PostV1AdminBannedPlayersSteamID(steamID uint64) (*http.Response, error) {
	return c.DoErr(http.MethodPost, fmt.Sprintf("/v1/admin/bannedPlayers/%d", steamID), jsonMap(nil), httpHeaders(nil))
}

// DeleteV1AdminBannedPlayersSteamID unbans a player with the specified Steam ID and returns the HTTP response.
//
//	DELETE /v1/admin/bannedPlayers/{steam_id}
func (c *HTTPClient) DeleteV1AdminBannedPlayersSteamID(steamID uint64) (*http.Response, error) {
	return c.DoErr(http.MethodDelete, fmt.Sprintf("/v1/admin/bannedPlayers/%d", steamID), jsonMap(nil), httpHeaders(nil))
}

// region /v1/admin/kickedPlayers
// todo: v1/admin/kickedPlayers endpoints
