package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GcbResourceType struct {

	// 功能说明：实例类型。
	ResourceType string `json:"resource_type"`
}

func (o GcbResourceType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GcbResourceType struct{}"
	}

	return strings.Join([]string{"GcbResourceType", string(data)}, " ")
}
