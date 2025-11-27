package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddonObjectMeta struct {

	// 唯一标识符
	Uid *string `json:"uid,omitempty"`

	// 对象的名称
	Name *string `json:"name,omitempty"`

	// 对象的标签
	Labels map[string]string `json:"labels,omitempty"`

	// 对象的注解
	Annotations map[string]string `json:"annotations,omitempty"`

	// 创建时间
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// 更新时间
	UpdateTimestamp *string `json:"updateTimestamp,omitempty"`
}

func (o AddonObjectMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddonObjectMeta struct{}"
	}

	return strings.Join([]string{"AddonObjectMeta", string(data)}, " ")
}
