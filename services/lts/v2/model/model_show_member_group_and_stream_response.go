package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMemberGroupAndStreamResponse Response Object
type ShowMemberGroupAndStreamResponse struct {
	Results        *[]MemberGroupandStreamResults `json:"results,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ShowMemberGroupAndStreamResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMemberGroupAndStreamResponse struct{}"
	}

	return strings.Join([]string{"ShowMemberGroupAndStreamResponse", string(data)}, " ")
}
