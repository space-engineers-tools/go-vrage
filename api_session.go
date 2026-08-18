package vrage

import (
	"log"
	"time"
)

// APISession provides access to the /v1/session API routes.
type APISession struct {
	http *HTTPClient
}

// region /v1/session
// todo

// region /v1/session/asteroids
// todo

// region /v1/session/chat

// APISessionPlayer represents a player in the session.
type APISessionPlayer struct {
	SteamID      uint64 `json:"SteamID"`
	DisplayName  string `json:"DisplayName"`
	FactionName  string `json:"FactionName"`
	FactionTag   string `json:"FactionTag"`
	PromoteLevel int    `json:"PromoteLevel"`
	Ping         int    `json:"Ping"`
}

// APISessionPlayersData represents the data returned by the GET /v1/session/players endpoint.
type APISessionPlayersData struct {
	Players []APISessionPlayer `json:"Players"`
}

// Players fetches the list of current players in the session.
func (s *APISession) Players() (APIResponseWithData[APISessionPlayersData], error) {
	var responseStruct APIResponseWithData[APISessionPlayersData]
	httpResponse, err := s.http.GetV1SessionPlayers()
	if err != nil {
		return responseStruct, err
	}
	defer func() {
		if err := httpResponse.Body.Close(); err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}()

	responseStruct, err = parseResponse[APIResponseWithData[APISessionPlayersData]](httpResponse)
	if err != nil {
		return responseStruct, err
	}

	return responseStruct, nil
}

// APISessionChatMessage represents a single chat message.
type APISessionChatMessage struct {
	SteamID     uint64 `json:"SteamID"`
	DisplayName string `json:"DisplayName"`
	Content     string `json:"Content"`
	// Timestamp is returned by the API as a string.
	Timestamp string `json:"Timestamp"`
}

// TimestampAsTime converts the Timestamp string to a time.Time object.
func (m *APISessionChatMessage) TimestampAsTime() (time.Time, error) {
	return DotNetTimestampToTime(m.Timestamp)
}

// APISessionChatData represents the data returned by the GET /v1/session/chat endpoint.
type APISessionChatData struct {
	Messages []APISessionChatMessage `json:"Messages"`
}

// ChatGet fetches chat messages.
func (s *APISession) ChatGet() (APIResponseWithData[APISessionChatData], error) {
	var responseStruct APIResponseWithData[APISessionChatData]
	httpResponse, err := s.http.GetV1SessionChat()
	if err != nil {
		return responseStruct, err
	}
	defer func() {
		if err := httpResponse.Body.Close(); err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}()

	responseStruct, err = parseResponse[APIResponseWithData[APISessionChatData]](httpResponse)
	if err != nil {
		return responseStruct, err
	}

	return responseStruct, nil
}

// ChatSend sends a chat message.
func (s *APISession) ChatSend(message string) (APIResponseWithoutData, error) {
	var responseStruct APIResponseWithoutData
	httpResponse, err := s.http.PostV1SessionChat(message)
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

// region /v1/session/floatingObjects
// todo

// region /v1/session/grids
// todo

// region /v1/session/planets
// todo

// region /v1/session/players
// todo

// region /v1/session/poweredGrids
// todo
