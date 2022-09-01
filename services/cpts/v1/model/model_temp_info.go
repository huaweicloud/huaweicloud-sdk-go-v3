package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TempInfo struct {

	// id
	Id *int32 `json:"id,omitempty" xml:"id"`

	// project_id
	ProjectId *int32 `json:"project_id,omitempty" xml:"project_id"`

	// name
	Name *string `json:"name,omitempty" xml:"name"`

	// description
	Description *string `json:"description,omitempty" xml:"description"`

	// variables
	Variables *string `json:"variables,omitempty" xml:"variables"`

	// contents
	Contents *[]interface{} `json:"contents,omitempty" xml:"contents"`

	// temp_type
	TempType *int32 `json:"temp_type,omitempty" xml:"temp_type"`

	// for_loop_params
	ForLoopParams *[]interface{} `json:"for_loop_params,omitempty" xml:"for_loop_params"`

	LogicController *LogicController `json:"logic_controller,omitempty" xml:"logic_controller"`

	// enable_pre
	EnablePre *bool `json:"enable_pre,omitempty" xml:"enable_pre"`
}

func (o TempInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TempInfo struct{}"
	}

	return strings.Join([]string{"TempInfo", string(data)}, " ")
}
