package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ShowFederationResponse Response Object
type ShowFederationResponse struct {

	// **参数解释**：Ray联邦端点ID。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。
	Id *string `json:"id,omitempty"`

	// **参数解释**：Ray联邦端点名称。 **取值范围**：支持小写字母、数字、中划线，必须字母开头、字母或数字结尾，长度为1~63个字符。
	Name *string `json:"name,omitempty"`

	// **参数解释**：Ray联邦端点显示名称。 **取值范围**：长度为1~63个字符，支持字母、数字、中划线、汉字。
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释**：Ray联邦端点描述。 **取值范围**：不允许输入感叹号、小于号、大于号、等号、与号、双引号、单引号、艾特符、井号、百分号、分号、左括号、右括号、竖线、星号、反斜杠。长度为0~1024个字符。
	Description *string `json:"description,omitempty"`

	// **参数解释**：Ray联邦端点状态。 **取值范围**：   - RUNNING：运行中。   - LIMITED_AVAILABLE：部分端点可用。   - FAULT：故障。   - WAITING：等待中。
	Status *ShowFederationResponseStatus `json:"status,omitempty"`

	Error *ErrorMessage `json:"error,omitempty"`

	// **参数解释**：RayCluster端点信息列表。
	RayClusterList *[]RayFederationRayClusterResponse `json:"ray_cluster_list,omitempty"`

	// **参数解释**：创建时间。使用UTC时间格式，格式为yyyy-MM-ddTHH:mm:ssZ，例如2023-05-30T12:24:30.401Z。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// **参数解释**：更新时间。使用UTC时间格式，格式为yyyy-MM-ddTHH:mm:ssZ，例如2023-05-30T12:24:30.401Z。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	Creator        *User `json:"creator,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowFederationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFederationResponse struct{}"
	}

	return strings.Join([]string{"ShowFederationResponse", string(data)}, " ")
}

type ShowFederationResponseStatus struct {
	value string
}

type ShowFederationResponseStatusEnum struct {
	RUNNING           ShowFederationResponseStatus
	LIMITED_AVAILABLE ShowFederationResponseStatus
	FAULT             ShowFederationResponseStatus
	WAITING           ShowFederationResponseStatus
}

func GetShowFederationResponseStatusEnum() ShowFederationResponseStatusEnum {
	return ShowFederationResponseStatusEnum{
		RUNNING: ShowFederationResponseStatus{
			value: "RUNNING",
		},
		LIMITED_AVAILABLE: ShowFederationResponseStatus{
			value: "LIMITED_AVAILABLE",
		},
		FAULT: ShowFederationResponseStatus{
			value: "FAULT",
		},
		WAITING: ShowFederationResponseStatus{
			value: "WAITING",
		},
	}
}

func (c ShowFederationResponseStatus) Value() string {
	return c.value
}

func (c ShowFederationResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFederationResponseStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
