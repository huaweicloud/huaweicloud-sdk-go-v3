package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// UpdateV2RayFederationResponse Response Object
type UpdateV2RayFederationResponse struct {

	// **参数解释**：Ray联邦端点ID。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。
	Id *string `json:"id,omitempty"`

	// **参数解释**：Ray联邦端点名称。 **取值范围**：支持小写字母、数字、中划线，必须字母开头、字母或数字结尾，长度为1~63个字符。
	Name *string `json:"name,omitempty"`

	// **参数解释**：Ray联邦端点显示名称。 **取值范围**：长度为1~63个字符，支持字母、数字、中划线、汉字。
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释**：Ray联邦端点描述。 **取值范围**：不允许输入感叹号、小于号、大于号、等号、与号、双引号、单引号。字符长度为0~1024个字符。
	Description *string `json:"description,omitempty"`

	// **参数解释**：Ray联邦端点状态。 **取值范围**：   - RUNNING：运行中。   - LIMITED_AVAILABLE：部分端点可用。   - FAULT：故障。   - WAITING：等待中。
	Status *UpdateV2RayFederationResponseStatus `json:"status,omitempty"`

	Error *Error `json:"error,omitempty"`

	// **参数解释**：RayCluster 端点信息列表。
	RayClusterList *[]V2RayFederationRayClusterResponse `json:"ray_cluster_list,omitempty"`

	// **参数解释**：资源创建时间。 **取值范围**：ISO 8601格式，yyyy-MM-dd'T'HH:mm:ss.SSSZ，UTC时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// **参数解释**：资源更新时间。 **取值范围**：ISO 8601格式，yyyy-MM-dd'T'HH:mm:ss.SSSZ，UTC时间。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	Creator        *Creator `json:"creator,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o UpdateV2RayFederationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateV2RayFederationResponse struct{}"
	}

	return strings.Join([]string{"UpdateV2RayFederationResponse", string(data)}, " ")
}

type UpdateV2RayFederationResponseStatus struct {
	value string
}

type UpdateV2RayFederationResponseStatusEnum struct {
	RUNNING           UpdateV2RayFederationResponseStatus
	LIMITED_AVAILABLE UpdateV2RayFederationResponseStatus
	FAULT             UpdateV2RayFederationResponseStatus
	WAITING           UpdateV2RayFederationResponseStatus
}

func GetUpdateV2RayFederationResponseStatusEnum() UpdateV2RayFederationResponseStatusEnum {
	return UpdateV2RayFederationResponseStatusEnum{
		RUNNING: UpdateV2RayFederationResponseStatus{
			value: "RUNNING",
		},
		LIMITED_AVAILABLE: UpdateV2RayFederationResponseStatus{
			value: "LIMITED_AVAILABLE",
		},
		FAULT: UpdateV2RayFederationResponseStatus{
			value: "FAULT",
		},
		WAITING: UpdateV2RayFederationResponseStatus{
			value: "WAITING",
		},
	}
}

func (c UpdateV2RayFederationResponseStatus) Value() string {
	return c.value
}

func (c UpdateV2RayFederationResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateV2RayFederationResponseStatus) UnmarshalJSON(b []byte) error {
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
