package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 获取预检查结果返回体
type QueryPreCheckResp struct {

	// 预检查id。
	PrecheckId *string `json:"precheck_id,omitempty" xml:"precheck_id"`

	// 返回的预检查结果是否通过。true表示预检查通过，通过后才可进行启动任务。
	Result *bool `json:"result,omitempty" xml:"result"`

	// 预检查进度百分比。
	Process *string `json:"process,omitempty" xml:"process"`

	// 预检查通过百分比。
	TotalPassedRate *string `json:"total_passed_rate,omitempty" xml:"total_passed_rate"`

	// RDS实例id。
	RdsInstanceId *string `json:"rds_instance_id,omitempty" xml:"rds_instance_id"`

	// 迁移方向
	JobDirection *QueryPreCheckRespJobDirection `json:"job_direction,omitempty" xml:"job_direction"`

	// 预检查各项结果。
	PrecheckResult *[]PrecheckResult `json:"precheck_result,omitempty" xml:"precheck_result"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg"`

	// 任务错误码。
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`
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
	UP_     QueryPreCheckRespJobDirection
	DOWN_   QueryPreCheckRespJobDirection
	NON_DBS QueryPreCheckRespJobDirection
}

func GetQueryPreCheckRespJobDirectionEnum() QueryPreCheckRespJobDirectionEnum {
	return QueryPreCheckRespJobDirectionEnum{
		UP_: QueryPreCheckRespJobDirection{
			value: "up-入云 灾备场景时对应本云为备",
		},
		DOWN_: QueryPreCheckRespJobDirection{
			value: "down-出云 灾备场景时对应本云为主",
		},
		NON_DBS: QueryPreCheckRespJobDirection{
			value: "non-dbs-自建",
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
