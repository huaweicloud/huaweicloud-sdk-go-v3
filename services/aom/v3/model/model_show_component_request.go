package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowComponentRequest Request Object
type ShowComponentRequest struct {

	// 组件id
	ComponentId string `json:"component_id"`
}

func (o ShowComponentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowComponentRequest struct{}"
	}

	return strings.Join([]string{"ShowComponentRequest", string(data)}, " ")
}
