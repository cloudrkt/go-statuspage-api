package statuspage

import (
	"errors"
	"fmt"
	"time"
)

type Subscriber struct {
	ID            *string    `json:"id,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	PhoneCountry  *string    `json:"phone_country,omitempty"`
	PhoneNumber   *string    `json:"phone_number,omitempty"`
	Email         *string    `json:"email,omitempty"`
	SkipNotify    *bool      `json:"skip_confirmation_notification,omitempty"`
	Mode          *string    `json:"mode,omitempty"`
	QuarantinedAt *time.Time `json:"quarantined_at,omitempty"`
	PurgeAt       *time.Time `json:"purge_at,omitempty"`
	Components    []*string  `json:"components,omitempty"`
}

type NewSubscriber struct {
	Email string `json:"email,omitempty"`
}

type SubscriberResponse []Subscriber

func (s *NewSubscriber) String() string {
	return encodeParams(map[string]interface{}{
		"subscriber[email]": s.Email,
	})
}

func (c *Client) GetAllSubscribers() ([]Subscriber, error) {
	return c.doGetSubscribers("subscribers.json")
}

func (c *Client) doGetSubscribers(path string) ([]Subscriber, error) {
	resp := SubscriberResponse{}
	err := c.doGet(path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *Subscriber) String() string {
	return fmt.Sprintf("ID: %s Created: %s Email: %s", *s.ID, *s.CreatedAt, *s.Email)
}

func (c *Client) CreateSubscriber(email string) (*Subscriber, error) {
	s := &NewSubscriber{email}
	resp := &Subscriber{}

	err := c.doPost("subscribers.json", s, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) DeleteSubscriber(subscriber *Subscriber) (*Subscriber, error) {
	path := "subscribers/" + *subscriber.ID + ".json"
	resp := &Subscriber{}
	err := c.doDelete(path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type Searchcriber struct {
	ID string `json:"id,omitempty"`
}

func (c *Client) SearchEmailSubscriber(email string) (*Subscriber, error) {
	path := "subscribers.json?q=" + email + "&state=all&limit=1"

	var resp []*Subscriber
	err := c.doGet(path, nil, &resp)
	if err != nil {
		return nil, err
	}
	if len(resp) != 1 { // Should return exactly 1 match
		return nil, errors.New("subscriber not found: " + email)
	}
	return resp[0], nil

}
