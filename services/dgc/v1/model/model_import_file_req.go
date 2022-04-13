package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type ImportFileReq struct {
	Path *string `json:"path,omitempty"`
	// 公共作业参数

	Params *interface{} `json:"params,omitempty"`

	SameNamePolicy *ImportFileReqSameNamePolicy `json:"sameNamePolicy,omitempty"`
	// 指定作业参数

	JobsParam *interface{} `json:"jobsParam,omitempty"`

	ExecuteUser *string `json:"executeUser,omitempty"`
}

func (o ImportFileReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportFileReq struct{}"
	}

	return strings.Join([]string{"ImportFileReq", string(data)}, " ")
}

type ImportFileReqSameNamePolicy struct {
	value string
}

type ImportFileReqSameNamePolicyEnum struct {
	SKIP      ImportFileReqSameNamePolicy
	OVERWRITE ImportFileReqSameNamePolicy
}

func GetImportFileReqSameNamePolicyEnum() ImportFileReqSameNamePolicyEnum {
	return ImportFileReqSameNamePolicyEnum{
		SKIP: ImportFileReqSameNamePolicy{
			value: "SKIP",
		},
		OVERWRITE: ImportFileReqSameNamePolicy{
			value: "OVERWRITE",
		},
	}
}

func (c ImportFileReqSameNamePolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImportFileReqSameNamePolicy) UnmarshalJSON(b []byte) error {
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
