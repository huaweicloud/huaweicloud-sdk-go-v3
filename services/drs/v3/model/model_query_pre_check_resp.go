package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// QueryPreCheckResp 获取预检查结果返回体
type QueryPreCheckResp struct {

	// **参数解释**：  查询预检查结果的任务ID，用于唯一标识一个预检查任务。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	PrecheckId *string `json:"precheck_id,omitempty"`

	// **参数解释**：  预检查结果是否通过，标识任务预检查的整体通过情况。  **约束限制**：  不涉及。  **取值范围**：  - true：预检查通过，通过后才可进行启动任务 - false：预检查未通过，需查看precheck_result中的具体失败项  **默认取值**：  不涉及
	Result *bool `json:"result,omitempty"`

	// **参数解释**：  预检查进度百分比，表示预检查任务的执行进度。  **约束限制**：  不涉及。  **取值范围**：  0%-100%。  **默认取值**：  不涉及
	Process *string `json:"process,omitempty"`

	// **参数解释**：  预检查通过百分比，表示预检查项中通过的占比。  **约束限制**：  不涉及。  **取值范围**：  0%-100%。  **默认取值**：  不涉及
	TotalPassedRate *string `json:"total_passed_rate,omitempty"`

	// **参数解释**：  RDS实例ID，用于标识关联的华为云RDS数据库实例。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	RdsInstanceId *string `json:"rds_instance_id,omitempty"`

	// **参数解释**：  任务的迁移方向。  **约束限制**：  不涉及。  **取值范围**：  - up：入云，灾备场景时对应本云为备 - down：出云，灾备场景时对应本云为主 - non-dbs：自建  **默认取值**：  不涉及
	JobDirection *QueryPreCheckRespJobDirection `json:"job_direction,omitempty"`

	// 预检查各项结果。
	PrecheckResult *[]PrecheckResult `json:"precheck_result,omitempty"`

	// **参数解释**：  错误详细信息，当任务执行失败时返回的具体错误描述。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	ErrorMsg *string `json:"error_msg,omitempty"`

	// **参数解释**：  错误码，当任务报错时返回的错误代码。  **约束限制**：  不涉及。  **取值范围**：  格式为DRS.XXXXXX。  **默认取值**：  不涉及
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
