package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProjectRequestBody
type CreateProjectRequestBody struct {

	// name
	Name string `json:"name" xml:"name"`

	// description
	Description *string `json:"description,omitempty" xml:"description"`
}

func (o CreateProjectRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectRequestBody struct{}"
	}

	return strings.Join([]string{"CreateProjectRequestBody", string(data)}, " ")
}
