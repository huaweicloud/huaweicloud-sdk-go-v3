package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 客户端应用与客户端配额绑定信息
type AppQuotaAppBinding struct {

	// 客户端配额编号
	AppQuotaId *string `json:"app_quota_id,omitempty" xml:"app_quota_id"`

	// 客户端应用编号
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	// 绑定时间
	BoundTime *sdktime.SdkTime `json:"bound_time,omitempty" xml:"bound_time"`
}

func (o AppQuotaAppBinding) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppQuotaAppBinding struct{}"
	}

	return strings.Join([]string{"AppQuotaAppBinding", string(data)}, " ")
}
