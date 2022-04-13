package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteSpecifiedConsumerGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSpecifiedConsumerGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSpecifiedConsumerGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteSpecifiedConsumerGroupResponse", string(data)}, " ")
}
