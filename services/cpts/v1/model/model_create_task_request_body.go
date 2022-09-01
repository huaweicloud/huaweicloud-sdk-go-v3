package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTaskRequestBody
type CreateTaskRequestBody struct {

	// name
	Name string `json:"name" xml:"name"`

	// project_id
	ProjectId int32 `json:"project_id" xml:"project_id"`

	// temps
	Temps *[]string `json:"temps,omitempty" xml:"temps"`

	// operate_mode
	OperateMode *int32 `json:"operate_mode,omitempty" xml:"operate_mode"`

	// bench_concurrent
	BenchConcurrent *int32 `json:"bench_concurrent,omitempty" xml:"bench_concurrent"`
}

func (o CreateTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTaskRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTaskRequestBody", string(data)}, " ")
}
