package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobScriptOrderListProp 脚本工单属性
type JobScriptOrderListProp struct {

	// CMDB服务实例区域id，可能有多个
	RegionIds *string `json:"region_ids,omitempty"`
}

func (o JobScriptOrderListProp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobScriptOrderListProp struct{}"
	}

	return strings.Join([]string{"JobScriptOrderListProp", string(data)}, " ")
}
