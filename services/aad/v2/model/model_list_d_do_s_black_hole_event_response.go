package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDDoSBlackHoleEventResponse Response Object
type ListDDoSBlackHoleEventResponse struct {

	// items
	Items          *[]ListBlackHoleEventRecordItem `json:"items,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListDDoSBlackHoleEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDDoSBlackHoleEventResponse struct{}"
	}

	return strings.Join([]string{"ListDDoSBlackHoleEventResponse", string(data)}, " ")
}
