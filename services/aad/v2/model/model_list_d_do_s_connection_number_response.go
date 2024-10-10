package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDDoSConnectionNumberResponse Response Object
type ListDDoSConnectionNumberResponse struct {

	// items
	Data           *[]ListConnectionNumberData `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListDDoSConnectionNumberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDDoSConnectionNumberResponse struct{}"
	}

	return strings.Join([]string{"ListDDoSConnectionNumberResponse", string(data)}, " ")
}
