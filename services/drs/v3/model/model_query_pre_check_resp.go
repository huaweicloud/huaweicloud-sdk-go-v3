package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// QueryPreCheckResp 获取预检查结果返回体
type QueryPreCheckResp struct {

	// 预检查id。
	PrecheckId *string `json:"precheck_id,omitempty"`

	// 返回的预检查结果是否通过。true表示预检查通过，通过后才可进行启动任务。
	Result *bool `json:"result,omitempty"`

	// 预检查进度百分比。
	Process *string `json:"process,omitempty"`

	// 预检查通过百分比。
	TotalPassedRate *string `json:"total_passed_rate,omitempty"`

	// RDS实例id。
	RdsInstanceId *string `json:"rds_instance_id,omitempty"`

	// 迁移方向。 - up-入云 灾备场景时对应本云为备 - down-出云 灾备场景时对应本云为主 - non-dbs-自建
	JobDirection *QueryPreCheckRespJobDirection `json:"job_direction,omitempty"`

	// 预检查各项结果。
	PrecheckResult *[]PrecheckResult `json:"precheck_result,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 任务错误码。
	ErrorCode *string `json:"error_code,omitempty"`
}

func (o QueryPreCheckResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryPreCheckResp struct{}"
	}

	return strings.Join([]string{"QueryPreCheckResp", string(data)}, " ")
}

type QueryPreCheckRespJobDirection struct {
	value string
}

type QueryPreCheckRespJobDirectionEnum struct {
	UP      QueryPreCheckRespJobDirection
	DOWN    QueryPreCheckRespJobDirection
	NON_DBS QueryPreCheckRespJobDirection
}

func GetQueryPreCheckRespJobDirectionEnum() QueryPreCheckRespJobDirectionEnum {
	return QueryPreCheckRespJobDirectionEnum{
		UP: QueryPreCheckRespJobDirection{
			value: "up",
		},
		DOWN: QueryPreCheckRespJobDirection{
			value: "down",
		},
		NON_DBS: QueryPreCheckRespJobDirection{
			value: "non-dbs",
		},
	}
}

func (c QueryPreCheckRespJobDirection) Value() string {
	return c.value
}

func (c QueryPreCheckRespJobDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryPreCheckRespJobDirection) UnmarshalJSON(b []byte) error {
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
