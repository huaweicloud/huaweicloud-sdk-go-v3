package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 诊断维度
type DiagnosisDimension struct {
	// 诊断维度名称

	Name DiagnosisDimensionName `json:"name"`
	// 诊断结果为异常的诊断项总数

	AbnormalNum int32 `json:"abnormal_num"`
	// 诊断失败的诊断项总数

	FailedNum int32 `json:"failed_num"`
	// 诊断项列表

	DiagnosisItemList []DiagnosisItem `json:"diagnosis_item_list"`
}

func (o DiagnosisDimension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiagnosisDimension struct{}"
	}

	return strings.Join([]string{"DiagnosisDimension", string(data)}, " ")
}

type DiagnosisDimensionName struct {
	value string
}

type DiagnosisDimensionNameEnum struct {
	NETWORK DiagnosisDimensionName
	STORAGE DiagnosisDimensionName
	LOAD    DiagnosisDimensionName
}

func GetDiagnosisDimensionNameEnum() DiagnosisDimensionNameEnum {
	return DiagnosisDimensionNameEnum{
		NETWORK: DiagnosisDimensionName{
			value: "network",
		},
		STORAGE: DiagnosisDimensionName{
			value: "storage",
		},
		LOAD: DiagnosisDimensionName{
			value: "load",
		},
	}
}

func (c DiagnosisDimensionName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DiagnosisDimensionName) UnmarshalJSON(b []byte) error {
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
