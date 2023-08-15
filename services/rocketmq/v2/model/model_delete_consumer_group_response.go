package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteConsumerGroupResponse Response Object
type DeleteConsumerGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteConsumerGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConsumerGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteConsumerGroupResponse", string(data)}, " ")
}
