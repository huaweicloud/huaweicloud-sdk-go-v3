package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchCreateOrDeleteQueueTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateOrDeleteQueueTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteQueueTagResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteQueueTagResponse", string(data)}, " ")
}
