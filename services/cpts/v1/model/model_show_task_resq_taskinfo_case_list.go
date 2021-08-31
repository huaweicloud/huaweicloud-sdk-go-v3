package model

import (
	"encoding/json"

	"strings"
)

type ShowTaskResqTaskinfoCaseList struct {
	// case_id

	CaseId *int32 `json:"case_id,omitempty"`
	// case_name

	CaseName *string `json:"case_name,omitempty"`
	// case_type

	CaseType *int32 `json:"case_type,omitempty"`
	// contents

	Contents *[]ShowTaskResqTaskinfoContents `json:"contents,omitempty"`
	// for_loop_params

	ForLoopParams *[]interface{} `json:"for_loop_params,omitempty"`
	// increase_setting

	IncreaseSetting *[]interface{} `json:"increase_setting,omitempty"`
	// stages

	Stages *[]interface{} `json:"stages,omitempty"`
	// status

	Status *int32 `json:"status,omitempty"`
	// temp_id

	TempId *int32 `json:"temp_id,omitempty"`
}

func (o ShowTaskResqTaskinfoCaseList) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowTaskResqTaskinfoCaseList struct{}"
	}

	return strings.Join([]string{"ShowTaskResqTaskinfoCaseList", string(data)}, " ")
}
