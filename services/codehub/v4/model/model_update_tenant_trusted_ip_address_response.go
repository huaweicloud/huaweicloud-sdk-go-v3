package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTenantTrustedIpAddressResponse Response Object
type UpdateTenantTrustedIpAddressResponse struct {

	// 关联结果
	Id *string `json:"id,omitempty"`

	// 用户id
	UserId *int32 `json:"user_id,omitempty"`

	// 租户id
	DomainId *string `json:"domain_id,omitempty"`

	// ip范围
	IpRange *string `json:"ip_range,omitempty"`

	// 格式类型，指定ip，ip范围，CIDR
	IpType *int32 `json:"ip_type,omitempty"`

	// 起始ip
	IpStart *string `json:"ip_start,omitempty"`

	// 结束ip
	IpEnd *string `json:"ip_end,omitempty"`

	// 是否允许访问代码仓库
	ViewFlag *int32 `json:"view_flag,omitempty"`

	// 是否允许下载代码
	DownloadFlag *int32 `json:"download_flag,omitempty"`

	// 是否允许提交代码
	UploadFlag *int32 `json:"upload_flag,omitempty"`

	// 备注
	Remark *string `json:"remark,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 排序
	OrderFlag      *int32 `json:"order_flag,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateTenantTrustedIpAddressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTenantTrustedIpAddressResponse struct{}"
	}

	return strings.Join([]string{"UpdateTenantTrustedIpAddressResponse", string(data)}, " ")
}
