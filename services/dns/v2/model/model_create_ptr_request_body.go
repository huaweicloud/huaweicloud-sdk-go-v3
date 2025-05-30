package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePtrRequestBody struct {
	Publicip *PublicIpReq `json:"publicip,omitempty"`

	// 反向解析记录对应的域名列表，最大个数不超过10个。
	Ptrdnames []string `json:"ptrdnames"`

	// 对反向解析记录的描述。
	Description *string `json:"description,omitempty"`

	// 反向解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。取值范围：1～2147483647
	Ttl *int32 `json:"ttl,omitempty"`

	// 资源标签。
	Tags *[]Tag `json:"tags,omitempty"`

	// 反向解析关联的企业项目ID，长度不超过36个字符。默认值为0。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreatePtrRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePtrRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePtrRequestBody", string(data)}, " ")
}
