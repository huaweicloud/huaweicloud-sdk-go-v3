package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteDevicesResponse Response Object
type BatchDeleteDevicesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteDevicesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDevicesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteDevicesResponse", string(data)}, " ")
}
