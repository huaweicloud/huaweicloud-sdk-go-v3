package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAppsResponse Response Object
type BatchDeleteAppsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteAppsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAppsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteAppsResponse", string(data)}, " ")
}
