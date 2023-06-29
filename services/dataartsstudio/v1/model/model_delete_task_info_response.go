package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTaskInfoResponse Response Object
type DeleteTaskInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTaskInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTaskInfoResponse struct{}"
	}

	return strings.Join([]string{"DeleteTaskInfoResponse", string(data)}, " ")
}
