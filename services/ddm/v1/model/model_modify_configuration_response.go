package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyConfigurationResponse Response Object
type ModifyConfigurationResponse struct {

	// 是否需要重启。
	NeedRestart    *bool `json:"need_restart,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ModifyConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ModifyConfigurationResponse", string(data)}, " ")
}
