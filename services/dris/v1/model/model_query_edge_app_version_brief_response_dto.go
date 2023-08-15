package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type QueryEdgeAppVersionBriefResponseDto struct {

	// **参数说明**：用户自定义应用唯一ID。
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
	State *QueryEdgeAppVersionBriefResponseDtoState `json:"state,omitempty"`

	// **参数说明**：发布时间。
	PublishTime *string `json:"publish_time,omitempty"`

	// **参数说明**：下线时间。
	OffShelfTime *string `json:"off_shelf_time,omitempty"`
}

func (o QueryEdgeAppVersionBriefResponseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryEdgeAppVersionBriefResponseDto struct{}"
	}

	return strings.Join([]string{"QueryEdgeAppVersionBriefResponseDto", string(data)}, " ")
}

type QueryEdgeAppVersionBriefResponseDtoState struct {
	value string
}

type QueryEdgeAppVersionBriefResponseDtoStateEnum struct {
	DRAFT     QueryEdgeAppVersionBriefResponseDtoState
	PUBLISHED QueryEdgeAppVersionBriefResponseDtoState
	OFF_SHELF QueryEdgeAppVersionBriefResponseDtoState
}

func GetQueryEdgeAppVersionBriefResponseDtoStateEnum() QueryEdgeAppVersionBriefResponseDtoStateEnum {
	return QueryEdgeAppVersionBriefResponseDtoStateEnum{
		DRAFT: QueryEdgeAppVersionBriefResponseDtoState{
			value: "DRAFT",
		},
		PUBLISHED: QueryEdgeAppVersionBriefResponseDtoState{
			value: "PUBLISHED",
		},
		OFF_SHELF: QueryEdgeAppVersionBriefResponseDtoState{
			value: "OFF_SHELF",
		},
	}
}

func (c QueryEdgeAppVersionBriefResponseDtoState) Value() string {
	return c.value
}

func (c QueryEdgeAppVersionBriefResponseDtoState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryEdgeAppVersionBriefResponseDtoState) UnmarshalJSON(b []byte) error {
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
