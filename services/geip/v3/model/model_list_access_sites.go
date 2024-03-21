package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAccessSites struct {

	// 接入点的ID
	Id *string `json:"id,omitempty"`

	// - 功能说明：接入点名称 - 取值范围：1-64，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）
	Name *string `json:"name,omitempty"`

	// pop站点托管的region(id)
	ProxyRegion *string `json:"proxy_region,omitempty"`

	// 边缘站点az
	IecAzCode *string `json:"iec_az_code,omitempty"`

	// 英文名称
	EnName *string `json:"en_name,omitempty"`

	// 中文名称
	CnName *string `json:"cn_name,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o ListAccessSites) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessSites struct{}"
	}

	return strings.Join([]string{"ListAccessSites", string(data)}, " ")
}
