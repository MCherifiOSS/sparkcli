package api

import (
	"errors"
	"github.com/tdeckers/sparkcli/util"
	"log"
)

type MessageService struct {
	Client *util.Client
}

type Message struct {
	Id            string `json:"id,omitempty"`
	RoomId        string `json:"roomId,omitempty"`
	Text          string `json:"text,omitempty"`
	Files         string `json:"files,omitempty"`
	ToPersonId    string `json:"toPersonId,omitempty"`
	ToPersonEmail string `json:"toPersonEmail,omitempty"`
	PersonId      string `json:"personId,omitempty"`
	PersonEmail   string `json:"personEmail,omitempty"`
	Created       string `json:"created,omitempty"`
}

type MessageItems struct {
	Items []Message `json:"items"`
}

func (m MessageService) list() (*[]Message, error) {
	log.Fatal("Not implemented")
	return nil, nil
}

func (m MessageService) List(roomId string) (*[]Message, error) {
	req, err := m.Client.NewGetRequest("/messages?roomId=" + roomId)
	if err != nil {
		return nil, err
	}
	var result MessageItems
	_, err = m.Client.Do(req, &result)
	if err != nil {
		return nil, err
	}
	return &result.Items, nil
}

// TODO: create different version, or update, to support direct msgs.
func (m MessageService) Create(roomId string, txt string) (*Message, error) {
	// Check for default roomId
	config := util.GetConfiguration()
	if roomId == "-" {
		if config.DefaultRoomId != "" {
			roomId = config.DefaultRoomId
		} else {
			return nil, errors.New("No DefaultRoomId configured.")
		}
	}

	msg := Message{RoomId: roomId, Text: txt}
	req, err := m.Client.NewPostRequest("/messages", msg)
	if err != nil {
		return nil, err
	}
	var result Message
	_, err = m.Client.Do(req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (m MessageService) Get(id string) (*Message, error) {
	if id == "" {
		return nil, errors.New("id can't be empty when getting message")
	}
	req, err := m.Client.NewGetRequest("/messages/" + id)
	if err != nil {
		return nil, err
	}
	var result Message
	_, err = m.Client.Do(req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (m MessageService) Delete(id string) error {
	if id == "" {
		return errors.New("id can't be empty when deleting a message")
	}
	req, err := m.Client.NewDeleteRequest("/messages/" + id)
	if err != nil {
		return err
	}
	_, err = m.Client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil //success
}
