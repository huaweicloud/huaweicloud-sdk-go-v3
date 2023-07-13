package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowComponentByNameRequest Request Object
type ShowComponentByNameRequest struct {

	// 应用id
	ApplicationId string `json:"application_id"`

	// 组件名称
	ComponentName string `json:"component_name"`
}

func (o ShowComponentByNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowComponentByNameRequest struct{}"
	}

	return strings.Join([]string{"ShowComponentByNameRequest", string(data)}, " ")
}
