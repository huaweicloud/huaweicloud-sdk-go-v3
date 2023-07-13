package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIpGroupIpOption create ip group ip option
type CreateIpGroupIpOption struct {

	// IP地址组中的IP网段，cidr格式。
	Cidr string `json:"cidr"`

	// IP地址组中的IP网段描述，取值范围：0~255个字符之间，禁止输入字符：<>。
	Description *string `json:"description,omitempty"`
}

func (o CreateIpGroupIpOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpGroupIpOption struct{}"
	}

	return strings.Join([]string{"CreateIpGroupIpOption", string(data)}, " ")
}
