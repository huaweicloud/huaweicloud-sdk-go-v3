package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowConsumerListOrDetailsResponse struct {

	// Topic列表。
	Topics         *[]string `json:"topics,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowConsumerListOrDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConsumerListOrDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowConsumerListOrDetailsResponse", string(data)}, " ")
}
