package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTempRequestBody
type UpdateTempRequestBody struct {

	// id
	Id int32 `json:"id" xml:"id"`

	// project_id
	ProjectId int32 `json:"project_id" xml:"project_id"`

	// name
	Name string `json:"name" xml:"name"`

	// temp_type
	TempType *int32 `json:"temp_type,omitempty" xml:"temp_type"`

	// description
	Description *string `json:"description,omitempty" xml:"description"`

	// for_loop_params
	ForLoopParams *[]interface{} `json:"for_loop_params,omitempty" xml:"for_loop_params"`

	// enable_pre
	EnablePre *bool `json:"enable_pre,omitempty" xml:"enable_pre"`

	// contents
	Contents *[]TempContentInfo `json:"contents,omitempty" xml:"contents"`
}

func (o UpdateTempRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTempRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateTempRequestBody", string(data)}, " ")
}
