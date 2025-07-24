package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIRackResponse Response Object
type ListIRackResponse struct {
	Iracks *[]IRack `json:"iracks,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListIRackResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIRackResponse struct{}"
	}

	return strings.Join([]string{"ListIRackResponse", string(data)}, " ")
}
