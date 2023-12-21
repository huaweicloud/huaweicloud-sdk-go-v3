package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateInstanceTagResponse Response Object
type BatchCreateInstanceTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateInstanceTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateInstanceTagResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateInstanceTagResponse", string(data)}, " ")
}
