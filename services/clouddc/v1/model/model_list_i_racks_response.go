package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIRacksResponse Response Object
type ListIRacksResponse struct {
	Iracks *[]IRack `json:"iracks,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListIRacksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIRacksResponse struct{}"
	}

	return strings.Join([]string{"ListIRacksResponse", string(data)}, " ")
}
