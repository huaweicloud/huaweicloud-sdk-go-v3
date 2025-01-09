package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchEnableAppsResponse Response Object
type BatchEnableAppsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchEnableAppsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchEnableAppsResponse struct{}"
	}

	return strings.Join([]string{"BatchEnableAppsResponse", string(data)}, " ")
}
