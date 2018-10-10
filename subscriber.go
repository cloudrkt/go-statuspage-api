package statuspage

import (
	"fmt"
	"time"
)

type Subscriber struct {
	ID            *string      `json:"id,omitempty"`
	CreatedAt     *time.Time   `json:"created_at,omitempty"`
	PhoneCountry  *string      `json:"phone_country,omitempty"`
	PhoneNumber   *string      `json:"phone_number,omitempty"`
	Email         *string      `json:"email,omitempty"`
	SkipNotify    *bool        `json:"skip_confirmation_notification,omitempty"`
	Mode          *string      `json:"mode,omitempty"`
	QuarantinedAt *time.Time   `json:"quarantined_at,omitempty"`
	PurgeAt       *time.Time   `json:"purge_at,omitempty"`
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
	var email, phonecountry, phonenumber, quarantineAt, purgeAt string

	if s.Email != nil {
		email = *s.Email
	}

	if s.PhoneCountry != nil {
		phonecountry = *s.PhoneCountry
	}

	if s.PhoneNumber != nil {
		phonenumber = *s.PhoneNumber
	}

	if s.QuarantinedAt != nil {
		quarantineAt = s.QuarantinedAt.String()
	}

	if s.PurgeAt != nil {
		purgeAt = s.PurgeAt.String()
	}

	line := "-----------------"
	out := fmt.Sprintf("\n%s\nID: %s\nCreated: %s\nEmail: %s\n PhoneNR: %s\n PhoneCountry: %s\n SkipNotify: %t\n Mode: %s\n QuarantinedAt: %s\n PurgeAt: %s\n%s",
		line,
		*s.ID,
		*s.CreatedAt,
		email,
		phonenumber,
		phonecountry,
		*s.SkipNotify,
		*s.Mode,
		quarantineAt,
		purgeAt,
		line,
	)
	return out
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
