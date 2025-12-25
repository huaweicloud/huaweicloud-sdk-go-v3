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

	// **参数解释**： 外部数据源类型。 **取值范围**： - OBS: obs数据源。 - LAKE_FORMATION: lake_formation数据源。 - MRS: mrs数据源。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 数据库。 **取值范围**： 不涉及。
	ConnectInfo *string `json:"connect_info,omitempty"`

	// **参数解释**： 用户名。 **取值范围**： 不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**： 版本。 **取值范围**： 不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释**： 配置状态。 **取值范围**： - 100: CREATING，创建中。 - 200: ACTIVE，已可用。 - 300: FAILED，已失败。 - 400: DELETED，已删除。 - 401: DELETING，删除中。 - 500: UPDATING，更新中。 - 600: PENDING_REBOOT，待重启。
	ConfigureStatus *string `json:"configure_status,omitempty"`

	// **参数解释**： 状态。 **取值范围**： - 100: CREATING, 创建中。 - 200: AVAILABLE, 可用。 - 300: FAILED, 失败。 - 303: CREATE_FAILED, 创建失败。 - 400: DELETED, 已删除。 - 304: DELETING, 删除中。 - 302: DELETE_FAILED, 删除失败。 - 800: FROZEN, 冻结。 - 801: POLICE_FROZEN, 警方冻结。 - 910: STOPPING, 停止中。 - 900: STOPPED, 已停止。 - 920: STARTING, 启动中。
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
