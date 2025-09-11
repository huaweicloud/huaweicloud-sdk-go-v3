package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerVersionResponse Response Object
type ShowServerVersionResponse struct {

	// 版本号信息
	Version        *string `json:"version,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowServerVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowServerVersionResponse", string(data)}, " ")
}
