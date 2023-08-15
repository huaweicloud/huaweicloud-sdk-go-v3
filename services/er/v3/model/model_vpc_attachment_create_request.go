package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VpcAttachmentCreateRequest VPC类型连接
type VpcAttachmentCreateRequest struct {

	// VPC的id，取值范围：最大长度36字节，带“-”连字符的UUID格式
	VpcId string `json:"vpc_id"`

	// VPC连接名字，取值范围：1-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）
	Name string `json:"name"`

	// VPC子网id，取值范围：最大长度36字节，带“-”连字符的UUID格式
	VirsubnetId string `json:"virsubnet_id"`

	// 描述信息，取值范围：0-255个字符，不能包含“<”和“>”
	Description *string `json:"description,omitempty"`

	// 默认为false，为true表示自动为vpc配置指向企业路由器的路由
	AutoCreateVpcRoutes *bool `json:"auto_create_vpc_routes,omitempty"`

	// 标签信息
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o VpcAttachmentCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcAttachmentCreateRequest struct{}"
	}

	return strings.Join([]string{"VpcAttachmentCreateRequest", string(data)}, " ")
}
