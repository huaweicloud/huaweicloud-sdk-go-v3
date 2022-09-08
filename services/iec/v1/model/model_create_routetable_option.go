package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建路由表请求体
type CreateRoutetableOption struct {

	// 路由表名称  取值范围：1-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）
	Name string `json:"name"`

	// 路由表所在的虚拟私有云ID
	VpcId string `json:"vpc_id"`

	// 路由表描述信息  取值范围：0-255个字符，不能包含“<”和“>”
	Description *string `json:"description,omitempty"`
}

func (o CreateRoutetableOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRoutetableOption struct{}"
	}

	return strings.Join([]string{"CreateRoutetableOption", string(data)}, " ")
}
