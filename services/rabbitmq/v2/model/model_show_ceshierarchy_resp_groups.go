package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowCeshierarchyRespGroups struct {

	// 消费组名称。
	Name *string `json:"name,omitempty"`
}

func (o ShowCeshierarchyRespGroups) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCeshierarchyRespGroups struct{}"
	}

	return strings.Join([]string{"ShowCeshierarchyRespGroups", string(data)}, " ")
}
