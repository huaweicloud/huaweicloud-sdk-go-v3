package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnusedIamAgencyDetails 未使用委托分析详细结果。
type UnusedIamAgencyDetails struct {

	// 用户使用委托的最后访问时间。
	LastAccessed *sdktime.SdkTime `json:"last_accessed,omitempty"`
}

func (o UnusedIamAgencyDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnusedIamAgencyDetails struct{}"
	}

	return strings.Join([]string{"UnusedIamAgencyDetails", string(data)}, " ")
}
