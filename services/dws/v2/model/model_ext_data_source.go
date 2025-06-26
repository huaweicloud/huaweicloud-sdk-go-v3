package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExtDataSource **参数解释**： 数据源信息。 **取值范围**： 不涉及。
type ExtDataSource struct {

	// **参数解释**： 数据源ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 数据源名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 类型。 **取值范围**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 数据库。 **取值范围**： 不涉及。
	ConnectInfo *string `json:"connect_info,omitempty"`

	// **参数解释**： 用户名。 **取值范围**： 不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**： 版本。 **取值范围**： 不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释**： 配置状态。 **取值范围**： 不涉及。
	ConfigureStatus *string `json:"configure_status,omitempty"`

	// **参数解释**： 状态。 **取值范围**： 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 外部数据源ID。 **取值范围**： 不涉及。
	DataSourceId *string `json:"data_source_id,omitempty"`

	// **参数解释**： 创建时间。 **取值范围**： 不涉及。
	Created *string `json:"created,omitempty"`

	// **参数解释**： 更新时间。 **取值范围**： 不涉及。
	Updated *string `json:"updated,omitempty"`

	// **参数解释**： 数据源更新时间。 **取值范围**： 不涉及。
	DataSourceUpdated *string `json:"data_source_updated,omitempty"`

	// **参数解释**： 扩展信息。 **取值范围**： 不涉及。
	ExtendProperties *interface{} `json:"extend_properties,omitempty"`

	// **参数解释**： 描述。 **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 失败原因。 **取值范围**： 不涉及。
	FailReason *string `json:"fail_reason,omitempty"`
}

func (o ExtDataSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtDataSource struct{}"
	}

	return strings.Join([]string{"ExtDataSource", string(data)}, " ")
}
