package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GcbResourceId struct {

	// 功能说明：实例ID。 取值范围：1-36个字符，支持数字、字母、_(下划线)、-（中划线）
	ResourceId string `json:"resource_id"`
}

func (o GcbResourceId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GcbResourceId struct{}"
	}

	return strings.Join([]string{"GcbResourceId", string(data)}, " ")
}
