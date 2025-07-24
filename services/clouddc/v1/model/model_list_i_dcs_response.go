package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIDcsResponse Response Object
type ListIDcsResponse struct {
	Idcs *[]IDc `json:"idcs,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListIDcsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIDcsResponse struct{}"
	}

	return strings.Join([]string{"ListIDcsResponse", string(data)}, " ")
}
