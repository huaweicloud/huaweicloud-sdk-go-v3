package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePtrRequestBody struct {

	// 反向解析记录对应的域名列表，最大个数不超过10个。
	Ptrdnames []string `json:"ptrdnames"`

	// 对反向解析记录的描述。
	Description *string `json:"description,omitempty"`

	// 反向解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。取值范围：1～2147483647
	Ttl *int32 `json:"ttl,omitempty"`
}

func (o UpdatePtrRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePtrRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePtrRequestBody", string(data)}, " ")
}
