package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Result struct {
	Message *string `json:"message,omitempty" xml:"message"`

	RowCount *int32 `json:"rowCount,omitempty" xml:"rowCount"`

	Rows *string `json:"rows,omitempty" xml:"rows"`

	Schema *string `json:"schema,omitempty" xml:"schema"`
}

func (o Result) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Result struct{}"
	}

	return strings.Join([]string{"Result", string(data)}, " ")
}
