package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceConsumerGroupResponse Response Object
type DeleteInstanceConsumerGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteInstanceConsumerGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceConsumerGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstanceConsumerGroupResponse", string(data)}, " ")
}
