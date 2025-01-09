package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstallAppResponse Response Object
type InstallAppResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o InstallAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstallAppResponse struct{}"
	}

	return strings.Join([]string{"InstallAppResponse", string(data)}, " ")
}
