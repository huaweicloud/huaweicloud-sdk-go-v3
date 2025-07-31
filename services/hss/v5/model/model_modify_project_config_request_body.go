package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyProjectConfigRequestBody 修改配置信息请求体
type ModifyProjectConfigRequestBody struct {

	// 配置信息列表
	ConfigInfoList []ProjectConfigInfo `json:"config_info_list"`

	// 是否级联修改。若配置为true且enterprise_project_id=all_granted_eps时，对所有企业项目进行修改；配置为false且enterprise_project_id=all_granted_eps时，仅对all_granted_eps修改；其他场景该字段不生效。
	Cascade *bool `json:"cascade,omitempty"`
}

func (o ModifyProjectConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyProjectConfigRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyProjectConfigRequestBody", string(data)}, " ")
}
