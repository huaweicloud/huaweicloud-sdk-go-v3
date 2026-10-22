package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// V2RayFederationInfo **参数解释**：Ray联邦端点详细信息。
type V2RayFederationInfo struct {

	// **参数解释**：Ray联邦端点ID。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。
	Id *string `json:"id,omitempty"`

	// **参数解释**：Ray联邦端点名称。 **取值范围**：支持小写字母、数字、中划线，必须字母开头、字母或数字结尾，长度为1~63个字符。
	Name *string `json:"name,omitempty"`

	// **参数解释**：Ray联邦端点显示名称。 **取值范围**：长度为1~63个字符，支持字母、数字、中划线、汉字。
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释**：Ray联邦端点描述。 **取值范围**：不允许输入感叹号、小于号、大于号、等号、与号、双引号、单引号。字符长度为0~1024个字符。
	Description *string `json:"description,omitempty"`

	// **参数解释**：Ray联邦端点状态。 **取值范围**：   - RUNNING：运行中。   - LIMITED_AVAILABLE：部分端点可用。   - FAULT：故障。   - WAITING：等待中。
	Status *V2RayFederationInfoStatus `json:"status,omitempty"`

	Error *Error `json:"error,omitempty"`

	// **参数解释**：RayCluster 端点信息列表。
	RayClusterList *[]V2RayFederationRayClusterResponse `json:"ray_cluster_list,omitempty"`

	// **参数解释**：资源创建时间。 **取值范围**：ISO 8601格式，yyyy-MM-dd'T'HH:mm:ss.SSSZ，UTC时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// **参数解释**：资源更新时间。 **取值范围**：ISO 8601格式，yyyy-MM-dd'T'HH:mm:ss.SSSZ，UTC时间。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	Creator *Creator `json:"creator,omitempty"`
}

func (o V2RayFederationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2RayFederationInfo struct{}"
	}

	return strings.Join([]string{"V2RayFederationInfo", string(data)}, " ")
}

type V2RayFederationInfoStatus struct {
	value string
}

type V2RayFederationInfoStatusEnum struct {
	RUNNING           V2RayFederationInfoStatus
	LIMITED_AVAILABLE V2RayFederationInfoStatus
	FAULT             V2RayFederationInfoStatus
	WAITING           V2RayFederationInfoStatus
}

func GetV2RayFederationInfoStatusEnum() V2RayFederationInfoStatusEnum {
	return V2RayFederationInfoStatusEnum{
		RUNNING: V2RayFederationInfoStatus{
			value: "RUNNING",
		},
		LIMITED_AVAILABLE: V2RayFederationInfoStatus{
			value: "LIMITED_AVAILABLE",
		},
		FAULT: V2RayFederationInfoStatus{
			value: "FAULT",
		},
		WAITING: V2RayFederationInfoStatus{
			value: "WAITING",
		},
	}
}

func (c V2RayFederationInfoStatus) Value() string {
	return c.value
}

func (c V2RayFederationInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *V2RayFederationInfoStatus) UnmarshalJSON(b []byte) error {
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
