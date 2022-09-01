package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CaseInfo struct {

	// case_id
	CaseId *int32 `json:"case_id,omitempty" xml:"case_id"`

	// case_name
	CaseName *string `json:"case_name,omitempty" xml:"case_name"`

	// case_type
	CaseType *int32 `json:"case_type,omitempty" xml:"case_type"`

	// contents
	Contents *[]Contents `json:"contents,omitempty" xml:"contents"`

	// for_loop_params
	ForLoopParams *[]interface{} `json:"for_loop_params,omitempty" xml:"for_loop_params"`

	// increase_setting
	IncreaseSetting *[]interface{} `json:"increase_setting,omitempty" xml:"increase_setting"`

	// stages
	Stages *[]interface{} `json:"stages,omitempty" xml:"stages"`

	// status
	Status *int32 `json:"status,omitempty" xml:"status"`

	// temp_id
	TempId *int32 `json:"temp_id,omitempty" xml:"temp_id"`

	// sort
	Sort *int32 `json:"sort,omitempty" xml:"sort"`
}

func (o CaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CaseInfo struct{}"
	}

	return strings.Join([]string{"CaseInfo", string(data)}, " ")
}
