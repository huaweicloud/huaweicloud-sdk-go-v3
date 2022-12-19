package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type CreateEdgeApplicationVersionResponse struct {

	// **参数说明**：用户自定义应用唯一ID。  **取值范围**：只允许字母、数字、下划线（_）、连接符（-）、美元符号（$）的组合。
	EdgeAppId *string `json:"edge_app_id,omitempty"`

	// **参数说明**：应用版本。
	Version *string `json:"version,omitempty"`

	// **参数说明**：应用描述。  **取值范围**：只允许中文、字母、数字、下划线（_）、中文分号（；）、中文冒号（：）、中文问号（？）、中文感叹号（！）中文逗号（，）、中文句号（。）、英文引号（;）、英文冒号（:）、英文逗号（,）、英文句号（.）、英文问号（?）、英文感叹号（!）、顿号（、）、连接符（-）的组合。
	Description *string `json:"description,omitempty"`

	// **参数说明**：创建时间。
	CreatedTime *string `json:"created_time,omitempty"`

	// **参数说明**：最后一次修改时间。
	LastModifiedTime *string `json:"last_modified_time,omitempty"`

	// **参数说明**：应用版本状态。  **取值范围**：  - DRAFT：草稿  - PUBLISHED：发布  - OFF_SHELF：下线
	State *CreateEdgeApplicationVersionResponseState `json:"state,omitempty"`

	// **参数说明**：启动命令。
	Command *[]string `json:"command,omitempty"`

	// **参数说明**：启动参数。
	Args *[]string `json:"args,omitempty"`

	ContainerSettings *ContainerSettingsDto `json:"container_settings,omitempty"`

	// **参数说明**：发布时间。
	PublishTime *string `json:"publish_time,omitempty"`

	// **参数说明**：下线时间。
	OffShelfTime   *string `json:"off_shelf_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEdgeApplicationVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeApplicationVersionResponse struct{}"
	}

	return strings.Join([]string{"CreateEdgeApplicationVersionResponse", string(data)}, " ")
}

type CreateEdgeApplicationVersionResponseState struct {
	value string
}

type CreateEdgeApplicationVersionResponseStateEnum struct {
	DRAFT     CreateEdgeApplicationVersionResponseState
	PUBLISHED CreateEdgeApplicationVersionResponseState
	OFF_SHELF CreateEdgeApplicationVersionResponseState
}

func GetCreateEdgeApplicationVersionResponseStateEnum() CreateEdgeApplicationVersionResponseStateEnum {
	return CreateEdgeApplicationVersionResponseStateEnum{
		DRAFT: CreateEdgeApplicationVersionResponseState{
			value: "DRAFT",
		},
		PUBLISHED: CreateEdgeApplicationVersionResponseState{
			value: "PUBLISHED",
		},
		OFF_SHELF: CreateEdgeApplicationVersionResponseState{
			value: "OFF_SHELF",
		},
	}
}

func (c CreateEdgeApplicationVersionResponseState) Value() string {
	return c.value
}

func (c CreateEdgeApplicationVersionResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEdgeApplicationVersionResponseState) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
