package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddTrustedIpAddressRequestBody 添加ip白名单请求体
type AddTrustedIpAddressRequestBody struct {

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
}

func (o AddTrustedIpAddressRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddTrustedIpAddressRequestBody struct{}"
	}

	return strings.Join([]string{"AddTrustedIpAddressRequestBody", string(data)}, " ")
}
