package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateV2EndpointResponse Response Object
type UpdateV2EndpointResponse struct {

	// **参数解释**：端点ID。 **取值范围**：长度为1~64个字符，支持大小写英文字母、数字、连字符。
	Id *string `json:"id,omitempty"`

	// **参数解释**：端点名称。 **取值范围**：支持小写字母、数字、中划线，必须字母开头、字母或数字结尾，长度为1~63个字符。
	Name *string `json:"name,omitempty"`

	// **参数解释**：端点显示名称。 **取值范围**：长度为1~63个字符，支持字母、数字、中划线、汉字。
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释**：资源描述信息。 **取值范围**：不允许输入感叹号、小于号、大于号、等号、与号、双引号、单引号、艾特符、井号、百分号、分号、左括号、右括号、竖线、星号、反斜杠。长度为0~512个字符。
	Description *string `json:"description,omitempty"`

	// **参数解释**：Endpoint状态。 **取值范围**：   - CREATING：创建中。   - RUNNING：运行中。   - CREATE_FAILED：创建失败。   - DELETING：删除中。   - DELETED：已删除。   - UPDATING：更新中。   - UPDATE_FAILED：更新失败。   - DELETE_FAILED：删除失败。   - FROZEN：冻结。   - TRIAL_ENDED：试用结束。   - ON_DEMAND_FROZEN：按需冻结。
	Status *string `json:"status,omitempty"`

	// **参数解释**：引擎配置。
	Engine map[string]interface{} `json:"engine,omitempty"`

	// **参数解释**：计算资源容量配置。
	Capacity *[]V2EndpointCapacityItem `json:"capacity,omitempty"`

	ResourceType *ResourceType `json:"resource_type,omitempty"`

	// **参数解释**：资源创建时间。 **取值范围**：ISO 8601格式，yyyy-MM-dd'T'HH:mm:ss.SSSZ，UTC时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// **参数解释**：资源更新时间。 **取值范围**：ISO 8601格式，yyyy-MM-dd'T'HH:mm:ss.SSSZ，UTC时间。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// **参数解释**：计算资源ID。 **取值范围**：长度为1~36个英文字母、数字和中划线的组合。
	ComputeId *string `json:"compute_id,omitempty"`

	Credential *CredentialConfig `json:"credential,omitempty"`

	Creator *Creator `json:"creator,omitempty"`

	// **参数解释**：资源标签列表。
	Tags *[]ResourceTag `json:"tags,omitempty"`

	Error *Error `json:"error,omitempty"`

	PrometheusConfig *PrometheusConfig `json:"prometheus_config,omitempty"`

	LtsConfig      *LtsConfig `json:"lts_config,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o UpdateV2EndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateV2EndpointResponse struct{}"
	}

	return strings.Join([]string{"UpdateV2EndpointResponse", string(data)}, " ")
}
