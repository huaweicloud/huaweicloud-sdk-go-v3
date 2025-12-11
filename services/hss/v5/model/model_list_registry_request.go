package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRegistryRequest Request Object
type ListRegistryRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 仓库名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	RegistryName *string `json:"registry_name,omitempty"`

	// **参数解释**: 仓库ID **取值范围**: 字符长度1-64位
	RegistryId *string `json:"registry_id,omitempty"`

	// **参数解释**: 镜像仓类型，不传默认返回所有类型。如果要查询多个类型，可以使用逗号分隔。 **取值范围**: - Harbor harbor - Jfrog jfrog - SwrPrivate swr私有 - SwrShared  swr共享 - SwrEnterprise  swr企业 - Other  其他仓库
	RegistryType *string `json:"registry_type,omitempty"`

	// **参数解释**: 镜像仓地址 **取值范围**: 字符长度1-256位
	RegistryAddr *string `json:"registry_addr,omitempty"`

	// **参数解释**： 接入状态 **约束限制**： 不涉及 **取值范围**： - success：接入成功 - fail：接入异常 - accessing：接入中  **默认取值**： 不涉及
	Status *string `json:"status,omitempty"`
}

func (o ListRegistryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRegistryRequest struct{}"
	}

	return strings.Join([]string{"ListRegistryRequest", string(data)}, " ")
}
