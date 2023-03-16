package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 设置项目级权限策略请求体
type UpdateDataPolicyReq struct {

	// 项目级删除策略（true：允许项目数据删除操作，false：不允许删除操作，默认为true）
	DataDelete bool `json:"data_delete"`

	// 项目级下载策略（true：允许项目数据下载操作，false：不允许下载操作，默认为true）
	DataDownload bool `json:"data_download"`

	// 项目级加密策略（true：允许项目数据加密操作，false：不允许加密操作，默认为false）
	DataEncrypted bool `json:"data_encrypted"`

	// 项目级分享策略（true：允许项目数据拷贝/引用操作，false：不允许拷贝/引用操作，默认为true）
	DataShare bool `json:"data_share"`
}

func (o UpdateDataPolicyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDataPolicyReq struct{}"
	}

	return strings.Join([]string{"UpdateDataPolicyReq", string(data)}, " ")
}
