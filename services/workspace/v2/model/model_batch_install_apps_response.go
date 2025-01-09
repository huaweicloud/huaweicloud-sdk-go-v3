package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchInstallAppsResponse Response Object
type BatchInstallAppsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchInstallAppsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchInstallAppsResponse struct{}"
	}

	return strings.Join([]string{"BatchInstallAppsResponse", string(data)}, " ")
}
