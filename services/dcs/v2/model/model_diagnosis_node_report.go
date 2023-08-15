package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DiagnosisNodeReport 节点诊断报告
type DiagnosisNodeReport struct {

	// 节点IP。例如：192.168.0.234:6379
	NodeIp string `json:"node_ip"`

	// 节点所在可用区Code
	AzCode string `json:"az_code"`

	// 节点所在分片的名称
	GroupName string `json:"group_name"`

	// 诊断结果为异常的诊断项总数
	AbnormalSum int32 `json:"abnormal_sum"`

	// 诊断失败的诊断项总数
	FailedSum int32 `json:"failed_sum"`

	// 节点角色
	Role DiagnosisNodeReportRole `json:"role"`

	// 诊断维度列表
	DiagnosisDimensionList []DiagnosisDimension `json:"diagnosis_dimension_list"`

	CommandTimeTakenList *CommandTimeTakenList `json:"command_time_taken_list"`
}

func (o DiagnosisNodeReport) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiagnosisNodeReport struct{}"
	}

	return strings.Join([]string{"DiagnosisNodeReport", string(data)}, " ")
}

type DiagnosisNodeReportRole struct {
	value string
}

type DiagnosisNodeReportRoleEnum struct {
	MASTER DiagnosisNodeReportRole
	SLAVE  DiagnosisNodeReportRole
}

func GetDiagnosisNodeReportRoleEnum() DiagnosisNodeReportRoleEnum {
	return DiagnosisNodeReportRoleEnum{
		MASTER: DiagnosisNodeReportRole{
			value: "master",
		},
		SLAVE: DiagnosisNodeReportRole{
			value: "slave",
		},
	}
}

func (c DiagnosisNodeReportRole) Value() string {
	return c.value
}

func (c DiagnosisNodeReportRole) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DiagnosisNodeReportRole) UnmarshalJSON(b []byte) error {
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
