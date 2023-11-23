package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLogConfigResponse Response Object
type UpdateLogConfigResponse struct {

	// 日志配置id
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateLogConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateLogConfigResponse", string(data)}, " ")
}
