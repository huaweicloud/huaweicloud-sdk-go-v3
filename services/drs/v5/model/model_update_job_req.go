package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateJobReq 更新指定ID任务请求体。
type UpdateJobReq struct {

	// 更新指定ID任务详情类型。  场景一：更新单个任务详情，取值： - name：更新该任务名称。 - description：更新该任务描述。 - re_create：配置中任务三天以后虚拟机删除后重建。 - expired_days：更新任务异常自动结束时间，单位为天。 - notify：更新任务异常通知信息。  场景二：更新批量异步任务详情，取值： - all：批量异步创建的任务，参数校验不通过，需要指定全部参数进行更新时。 - network：批量异步创建的任务，测试连接不通过，需要更新源库/目标库信息时。 - policy：批量异步创建的任务，需要更新任务配置时。 - db_object：批量异步创建的任务，需要更新对象信息时。 - precheck：批量异步创建的任务，需要重新预检查时。
	Type UpdateJobReqType `json:"type"`

	Params *UpdateJob `json:"params"`
}

func (o UpdateJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateJobReq struct{}"
	}

	return strings.Join([]string{"UpdateJobReq", string(data)}, " ")
}

type UpdateJobReqType struct {
	value string
}

type UpdateJobReqTypeEnum struct {
	NAME         UpdateJobReqType
	DESCRIPTION  UpdateJobReqType
	ALL          UpdateJobReqType
	NETWORK      UpdateJobReqType
	POLICY       UpdateJobReqType
	DB_OBJECT    UpdateJobReqType
	PRECHECK     UpdateJobReqType
	RE_CREATE    UpdateJobReqType
	EXPIRED_DAYS UpdateJobReqType
	NOTIFY       UpdateJobReqType
}

func GetUpdateJobReqTypeEnum() UpdateJobReqTypeEnum {
	return UpdateJobReqTypeEnum{
		NAME: UpdateJobReqType{
			value: "name",
		},
		DESCRIPTION: UpdateJobReqType{
			value: "description",
		},
		ALL: UpdateJobReqType{
			value: "all",
		},
		NETWORK: UpdateJobReqType{
			value: "network",
		},
		POLICY: UpdateJobReqType{
			value: "policy",
		},
		DB_OBJECT: UpdateJobReqType{
			value: "db_object",
		},
		PRECHECK: UpdateJobReqType{
			value: "precheck",
		},
		RE_CREATE: UpdateJobReqType{
			value: "re_create",
		},
		EXPIRED_DAYS: UpdateJobReqType{
			value: "expired_days",
		},
		NOTIFY: UpdateJobReqType{
			value: "notify",
		},
	}
}

func (c UpdateJobReqType) Value() string {
	return c.value
}

func (c UpdateJobReqType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateJobReqType) UnmarshalJSON(b []byte) error {
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
