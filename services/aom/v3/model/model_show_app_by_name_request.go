package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppByNameRequest Request Object
type ShowAppByNameRequest struct {

	// 应用唯一标识;字符集长度2-64，仅支持字符集：英文字母、数字、下划线、中划线、点；应用唯一标识与显示名称至少填写其一
	Name *string `json:"name,omitempty"`

	// 实体的显示名称；字符集长度2-64，仅支持字符集：中文字符、英文字母、数字、下划线、中划线、点；应用唯一标识与显示名称至少填写其一
	DisplayName *string `json:"display_name,omitempty"`
}

func (o ShowAppByNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppByNameRequest struct{}"
	}

	return strings.Join([]string{"ShowAppByNameRequest", string(data)}, " ")
}
