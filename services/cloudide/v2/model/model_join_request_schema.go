package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JoinRequestSchema the request body of join-request
type JoinRequestSchema struct {

	// the region of user
	Region *string `json:"region,omitempty"`

	// the name of user
	Name *string `json:"name,omitempty"`

	// the email of user
	Email *string `json:"email,omitempty"`

	// the organization of user
	Organization *string `json:"organization,omitempty"`

	// the phone_number of user
	PhoneNumber *string `json:"phone_number,omitempty"`

	// the invitation_code
	InvitationCode *string `json:"invitation_code,omitempty"`
}

func (o JoinRequestSchema) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JoinRequestSchema struct{}"
	}

	return strings.Join([]string{"JoinRequestSchema", string(data)}, " ")
}
