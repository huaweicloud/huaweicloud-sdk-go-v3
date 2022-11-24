package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDataPolicyResponse struct {

	// 项目级删除策略（true：允许项目数据删除操作，false：不允许删除操作，默认为true）
	DataDelete *bool `json:"data_delete,omitempty"`

	// 项目级下载策略（true：允许项目数据下载操作，false：不允许下载操作，默认为true）
	DataDownload *bool `json:"data_download,omitempty"`

	// 项目级加密策略（true：允许项目数据加密操作，false：不允许加密操作，默认为false）
	DataEncrypted *bool `json:"data_encrypted,omitempty"`

	// 项目级分享策略（true：允许项目数据拷贝/引用操作，false：不允许拷贝/引用操作，默认为true）
	DataShare      *bool `json:"data_share,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowDataPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowDataPolicyResponse", string(data)}, " ")
}
