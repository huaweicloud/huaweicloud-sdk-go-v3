package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemediationResourceParameter 合规规则修正执行的动态参数。
type RemediationResourceParameter struct {

	// 传入resourceID的参数名称。
	ResourceId string `json:"resource_id"`
}

func (o RemediationResourceParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemediationResourceParameter struct{}"
	}

	return strings.Join([]string{"RemediationResourceParameter", string(data)}, " ")
}
