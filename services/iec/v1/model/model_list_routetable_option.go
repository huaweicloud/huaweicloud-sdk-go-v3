package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 列表接口返回的路由表对象
type ListRoutetableOption struct {

	// 路由表ID  取值范围：标准UUID
	Id *string `json:"id,omitempty"`

	// 路由表名称  取值范围：0-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）
	Name *string `json:"name,omitempty"`

	// 路由表所在的虚拟私有云ID
	VpcId *string `json:"vpc_id,omitempty"`

	// 帐号ID
	DomainId *string `json:"domain_id,omitempty"`

	// 路由表描述信息  取值范围：0-255个字符，不能包含“<”和“>”
	Description *string `json:"description,omitempty"`

	// 是否为默认路由表  取值范围：true表示默认路由表；false表示自定义路由表
	Default *bool `json:"default,omitempty"`
}

func (o ListRoutetableOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRoutetableOption struct{}"
	}

	return strings.Join([]string{"ListRoutetableOption", string(data)}, " ")
}
