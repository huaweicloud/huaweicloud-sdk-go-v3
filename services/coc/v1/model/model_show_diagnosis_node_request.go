package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDiagnosisNodeRequest Request Object
type ShowDiagnosisNodeRequest struct {

	// 诊断工单ID
	TaskId string `json:"task_id"`

	// 诊断步骤编码
	Code ShowDiagnosisNodeRequestCode `json:"code"`

	// 被诊断实例ID
	InstanceId string `json:"instance_id"`
}

func (o ShowDiagnosisNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDiagnosisNodeRequest struct{}"
	}

	return strings.Join([]string{"ShowDiagnosisNodeRequest", string(data)}, " ")
}

type ShowDiagnosisNodeRequestCode struct {
	value string
}

type ShowDiagnosisNodeRequestCodeEnum struct {
	HOLMES_INSTALL    ShowDiagnosisNodeRequestCode
	DATA_COLLECTION   ShowDiagnosisNodeRequestCode
	DIAGNOSIS_FAULT   ShowDiagnosisNodeRequestCode
	HOLMES_UN_INSTALL ShowDiagnosisNodeRequestCode
	RDS_DIAGNOSIS     ShowDiagnosisNodeRequestCode
	DCS_DIAGNOSIS     ShowDiagnosisNodeRequestCode
	DMS_DIAGNOSIS     ShowDiagnosisNodeRequestCode
	ELB_DIAGNOSIS     ShowDiagnosisNodeRequestCode
}

func GetShowDiagnosisNodeRequestCodeEnum() ShowDiagnosisNodeRequestCodeEnum {
	return ShowDiagnosisNodeRequestCodeEnum{
		HOLMES_INSTALL: ShowDiagnosisNodeRequestCode{
			value: "holmesInstall",
		},
		DATA_COLLECTION: ShowDiagnosisNodeRequestCode{
			value: "dataCollection",
		},
		DIAGNOSIS_FAULT: ShowDiagnosisNodeRequestCode{
			value: "diagnosisFault",
		},
		HOLMES_UN_INSTALL: ShowDiagnosisNodeRequestCode{
			value: "holmesUnInstall",
		},
		RDS_DIAGNOSIS: ShowDiagnosisNodeRequestCode{
			value: "rdsDiagnosis",
		},
		DCS_DIAGNOSIS: ShowDiagnosisNodeRequestCode{
			value: "dcsDiagnosis",
		},
		DMS_DIAGNOSIS: ShowDiagnosisNodeRequestCode{
			value: "dmsDiagnosis",
		},
		ELB_DIAGNOSIS: ShowDiagnosisNodeRequestCode{
			value: "elbDiagnosis",
		},
	}
}

func (c ShowDiagnosisNodeRequestCode) Value() string {
	return c.value
}

func (c ShowDiagnosisNodeRequestCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDiagnosisNodeRequestCode) UnmarshalJSON(b []byte) error {
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
