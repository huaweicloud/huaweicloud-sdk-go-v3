package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOperationLogsByVaultNameRequest Request Object
type ListOperationLogsByVaultNameRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 恢复状态，包含如下：   - success : 成功   - skipped : 跳过   - failed : 失败   - running : 正在运行   - timeout : 超时   - waiting : 等待
	Status *string `json:"status,omitempty"`

	// 服务器名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 查询起始点
	Offset int32 `json:"offset"`

	// 每页显示个数
	Limit int32 `json:"limit"`

	// 查询时间范围
	LastDays *int32 `json:"last_days,omitempty"`
}

func (o ListOperationLogsByVaultNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOperationLogsByVaultNameRequest struct{}"
	}

	return strings.Join([]string{"ListOperationLogsByVaultNameRequest", string(data)}, " ")
}
