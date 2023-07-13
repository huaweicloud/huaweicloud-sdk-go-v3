package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIpGroupOption 创建IP地址组的详细信息。
type CreateIpGroupOption struct {

	// IP地址组名称，取值范围：1~64个字符之间，只能由数字、字母、中划线和中文组成。
	Name string `json:"name"`

	// IP地址组的描述信息，取值范围：0~255个字符之间，禁止输入字符：<>。
	Description *string `json:"description,omitempty"`

	// IP地址组中的IP网段列表，一次最多支持添加20个条目。
	IpList *[]CreateIpGroupIpOption `json:"ip_list,omitempty"`
}

func (o CreateIpGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpGroupOption struct{}"
	}

	return strings.Join([]string{"CreateIpGroupOption", string(data)}, " ")
}
