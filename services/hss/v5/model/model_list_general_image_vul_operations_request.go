package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGeneralImageVulOperationsRequest Request Object
type ListGeneralImageVulOperationsRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 漏洞id **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	VulId string `json:"vul_id"`

	// **参数解释**: 镜像id **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释**: 镜像类型 **约束限制**: 不涉及 **取值范围**: - local：本地镜像 - registry：仓库镜像  **默认取值**: 不涉及
	ImageType string `json:"image_type"`

	// 镜像名称
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**: 漏洞当前状态 **约束限制**: 不涉及 **取值范围**: - vul_status_unfix：未处理 - vul_status_ignored：已忽略  **默认取值**: 不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释**: 处理用户名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**: 操作类型 **约束限制**: 不涉及 **取值范围**: - ignore：忽略 - not_ignore:：取消忽略 - add_to_whitelist：加入白名单  **默认取值**: 不涉及
	HandleType *string `json:"handle_type,omitempty"`

	// **参数解释**: 应用名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	AppName *string `json:"app_name,omitempty"`

	// **参数解释**: 应用版本 **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	AppVersion *string `json:"app_version,omitempty"`

	// **参数解释**: 备注 **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	Remark *string `json:"remark,omitempty"`
}

func (o ListGeneralImageVulOperationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGeneralImageVulOperationsRequest struct{}"
	}

	return strings.Join([]string{"ListGeneralImageVulOperationsRequest", string(data)}, " ")
}
