package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunBaselineDetectRequest Request Object
type RunBaselineDetectRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 配置检测策略组ID **约束限制**: 不涉及 **取值范围**: 字符串大小0-128 **默认取值**: 不涉及
	GroupId string `json:"group_id"`
}

func (o RunBaselineDetectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunBaselineDetectRequest struct{}"
	}

	return strings.Join([]string{"RunBaselineDetectRequest", string(data)}, " ")
}
