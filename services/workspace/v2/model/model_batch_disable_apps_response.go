package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDisableAppsResponse Response Object
type BatchDisableAppsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDisableAppsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDisableAppsResponse struct{}"
	}

	return strings.Join([]string{"BatchDisableAppsResponse", string(data)}, " ")
}
