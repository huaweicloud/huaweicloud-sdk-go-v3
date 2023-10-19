package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppQuotaAppBinding 凭据与凭据配额绑定信息
type AppQuotaAppBinding struct {

	// 凭据配额编号
	AppQuotaId *string `json:"app_quota_id,omitempty"`

	// 凭据编号
	AppId *string `json:"app_id,omitempty"`

	// 绑定时间
	BoundTime *sdktime.SdkTime `json:"bound_time,omitempty"`
}

func (o AppQuotaAppBinding) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppQuotaAppBinding struct{}"
	}

	return strings.Join([]string{"AppQuotaAppBinding", string(data)}, " ")
}
