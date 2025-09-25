package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeSwitchesStatusRequestBody struct {

	// 是否开启
	Enabled *bool `json:"enabled,omitempty"`
}

func (o ChangeSwitchesStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSwitchesStatusRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeSwitchesStatusRequestBody", string(data)}, " ")
}
