package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDesignOperationResultRequest Request Object
type ShowDesignOperationResultRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 默认值：application/json;charset=UTF-8 可选，有Body体的情况下必选，没有Body体则无需填写和校验。
	ContentType *string `json:"Content-Type,omitempty"`

	// 批量操作类型。 枚举值：   - ER_REVERSE_DB: 关系建模逆向数据库   - TRANSFORM_LOGIC_MODEL: 逻辑模型转物理模型
	OperationType ShowDesignOperationResultRequestOperationType `json:"operation_type"`

	// 每页查询条数，即查询Y条数据。默认值50，取值范围[1,100]。
	Limit *int32 `json:"limit,omitempty"`

	// 查询起始坐标，即跳过X条数据，仅支持0或limit的整数倍，不满足则向下取整，默认值0。
	Offset *int32 `json:"offset,omitempty"`

	// 批量操作id，在逻辑模型转物理表时，填写的是逻辑模型的model_id，在逆向数据库时，填写的是目标模型的model_id。model_id可从接口[获取模型](ListWorkspaces.xml)中获取。
	OperationId *string `json:"operation_id,omitempty"`
}

func (o ShowDesignOperationResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDesignOperationResultRequest struct{}"
	}

	return strings.Join([]string{"ShowDesignOperationResultRequest", string(data)}, " ")
}

type ShowDesignOperationResultRequestOperationType struct {
	value string
}

type ShowDesignOperationResultRequestOperationTypeEnum struct {
	ER_REVERSE_DB         ShowDesignOperationResultRequestOperationType
	TRANSFORM_LOGIC_MODEL ShowDesignOperationResultRequestOperationType
}

func GetShowDesignOperationResultRequestOperationTypeEnum() ShowDesignOperationResultRequestOperationTypeEnum {
	return ShowDesignOperationResultRequestOperationTypeEnum{
		ER_REVERSE_DB: ShowDesignOperationResultRequestOperationType{
			value: "ER_REVERSE_DB",
		},
		TRANSFORM_LOGIC_MODEL: ShowDesignOperationResultRequestOperationType{
			value: "TRANSFORM_LOGIC_MODEL",
		},
	}
}

func (c ShowDesignOperationResultRequestOperationType) Value() string {
	return c.value
}

func (c ShowDesignOperationResultRequestOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDesignOperationResultRequestOperationType) UnmarshalJSON(b []byte) error {
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
