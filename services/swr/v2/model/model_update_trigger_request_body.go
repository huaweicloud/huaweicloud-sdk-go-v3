package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateTriggerRequestBody struct {
	// 是否生效,true启用，false不启用

	Enable string `json:"enable"`
}

func (o UpdateTriggerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTriggerRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateTriggerRequestBody", string(data)}, " ")
}
