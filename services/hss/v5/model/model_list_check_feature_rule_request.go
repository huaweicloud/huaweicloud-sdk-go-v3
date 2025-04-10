package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCheckFeatureRuleRequest Request Object
type ListCheckFeatureRuleRequest struct {

	// 企业项目ID，查询所有企业项目时填写：all_granted_eps
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 操作系统类型，包含如下2种。 - Linux - Windows
	OsType *string `json:"os_type,omitempty"`
}

func (o ListCheckFeatureRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCheckFeatureRuleRequest struct{}"
	}

	return strings.Join([]string{"ListCheckFeatureRuleRequest", string(data)}, " ")
}
