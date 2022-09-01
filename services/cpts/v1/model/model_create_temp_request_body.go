package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTempRequestBody
type CreateTempRequestBody struct {

	// project_id
	ProjectId int32 `json:"project_id" xml:"project_id"`

	// temp_type
	TempType int32 `json:"temp_type" xml:"temp_type"`

	// name
	Name string `json:"name" xml:"name"`

	// description
	Description *string `json:"description,omitempty" xml:"description"`

	// contents
	Contents *[]interface{} `json:"contents,omitempty" xml:"contents"`
}

func (o CreateTempRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTempRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTempRequestBody", string(data)}, " ")
}
