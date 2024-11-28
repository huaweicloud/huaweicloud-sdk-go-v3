package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnusedIamUserPasswordDetails 未使用的用户密码详情。
type UnusedIamUserPasswordDetails struct {

	// 用户密码的最后访问时间。
	LastAccessed *sdktime.SdkTime `json:"last_accessed,omitempty"`
}

func (o UnusedIamUserPasswordDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnusedIamUserPasswordDetails struct{}"
	}

	return strings.Join([]string{"UnusedIamUserPasswordDetails", string(data)}, " ")
}
