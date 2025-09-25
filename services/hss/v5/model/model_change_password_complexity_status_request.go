package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangePasswordComplexityStatusRequest Request Object
type ChangePasswordComplexityStatusRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 动作类型   - ignore ：对口令复杂度检测未通过的主机执行忽略动作   - unignore ：对已忽略的口令复杂度检测未通过的主机执行取消忽略动作
	Action string `json:"action"`

	Body *ChangePasswordComplexityStatusRequestBody `json:"body,omitempty"`
}

func (o ChangePasswordComplexityStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangePasswordComplexityStatusRequest struct{}"
	}

	return strings.Join([]string{"ChangePasswordComplexityStatusRequest", string(data)}, " ")
}
