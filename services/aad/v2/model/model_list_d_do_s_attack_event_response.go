package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDDoSAttackEventResponse Response Object
type ListDDoSAttackEventResponse struct {

	// total
	Total *int32 `json:"total,omitempty"`

	// data
	Data           *[]ListDDoSEventData `json:"data,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListDDoSAttackEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDDoSAttackEventResponse struct{}"
	}

	return strings.Join([]string{"ListDDoSAttackEventResponse", string(data)}, " ")
}
