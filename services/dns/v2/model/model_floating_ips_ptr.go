package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FloatingIpsPtr struct {

	// 反向解析记录对应的域名列表，最大个数不超过10个。
	Ptrdnames *[]string `json:"ptrdnames,omitempty"`

	// 反向解析记录的ID。
	Id *string `json:"id,omitempty"`

	Publicip *PublicIpRes `json:"publicip,omitempty"`

	// 对反向解析记录的描述。
	Description *string `json:"description,omitempty"`

	// 反向解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。
	Ttl *int32 `json:"ttl,omitempty"`

	// 资源状态。
	Status *string `json:"status,omitempty"`

	Links *PageLink `json:"links,omitempty"`

	// 反向解析关联的企业项目ID，长度不超过36个字符。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 资源标签。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o FloatingIpsPtr) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FloatingIpsPtr struct{}"
	}

	return strings.Join([]string{"FloatingIpsPtr", string(data)}, " ")
}
