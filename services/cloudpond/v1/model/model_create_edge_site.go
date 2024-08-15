package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateEdgeSite struct {

	// 边缘小站名称，最大支持长度为64个字节。只包含中文字符、英文字母（a-z、A-Z）、数字（0-9）、下划线（_）、中划线（-）
	Name string `json:"name"`

	// 边缘小站所属区域ID，最大长度为64个字节。只包含英文字母（a-z、A-Z）、数字（0-9）、下划线（_）、中划线（-）
	RegionId string `json:"region_id"`

	// 边缘小站描述，最大支持长度为255个字节，不允许包含<>
	Description *string `json:"description,omitempty"`

	Location *CreateLocation `json:"location"`

	// 企业项目Id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreateEdgeSite) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeSite struct{}"
	}

	return strings.Join([]string{"CreateEdgeSite", string(data)}, " ")
}
