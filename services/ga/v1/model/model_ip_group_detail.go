package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpGroupDetail IP地址组详情。
type IpGroupDetail struct {

	// IP地址组ID。
	Id *string `json:"id,omitempty"`

	// IP地址组名称。
	Name *string `json:"name,omitempty"`

	// IP地址组描述。
	Description *string `json:"description,omitempty"`

	Status *ConfigStatus `json:"status,omitempty"`

	// IP地址组中的IP网段列表。
	IpList *[]IpInfo `json:"ip_list,omitempty"`

	AssociatedListeners *[]ListenerAccessControlPolicy `json:"associated_listeners,omitempty"`

	// 创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o IpGroupDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpGroupDetail struct{}"
	}

	return strings.Join([]string{"IpGroupDetail", string(data)}, " ")
}
