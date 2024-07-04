package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteQueueInfoResponse Response Object
type DeleteQueueInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteQueueInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteQueueInfoResponse struct{}"
	}

	return strings.Join([]string{"DeleteQueueInfoResponse", string(data)}, " ")
}
