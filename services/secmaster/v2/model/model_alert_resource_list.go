package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlertResourceList struct {

	// 云服务资源id
	Id *string `json:"id,omitempty"`

	// 资源名称
	Name *string `json:"name,omitempty"`

	// 资源类型；引用华为云RMS type字段
	Type *string `json:"type,omitempty"`

	// 云服务名称；引用华为云RMS provider字段
	Provider *string `json:"provider,omitempty"`

	// 区域；按照华为云regionId填写，如cn-north-1等
	RegionId *string `json:"region_id,omitempty"`

	// 资源所属账号ID，UUID格式
	DomainId *string `json:"domain_id,omitempty"`

	// 资源所属项目ID，UUID格式
	ProjectId *string `json:"project_id,omitempty"`

	// 企业项目id
	EpId *string `json:"ep_id,omitempty"`

	// 企业项目名称
	EpName *string `json:"ep_name,omitempty"`

	// 资源标签 1、最多50个key/values对 2、values：最大255字符，取值范围：字母数字,空格,+, -, =, ., _, :, /,@
	Tags *string `json:"tags,omitempty"`
}

func (o AlertResourceList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertResourceList struct{}"
	}

	return strings.Join([]string{"AlertResourceList", string(data)}, " ")
}
