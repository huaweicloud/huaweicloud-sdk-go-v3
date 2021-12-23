package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowMemberDetailResponse struct {
	Member         *Member `json:"member,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMemberDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMemberDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowMemberDetailResponse", string(data)}, " ")
}
