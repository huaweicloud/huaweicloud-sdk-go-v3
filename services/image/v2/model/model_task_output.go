package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskOutput struct {
	Obs *TaskOutputObs `json:"obs"`
}

func (o TaskOutput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskOutput struct{}"
	}

	return strings.Join([]string{"TaskOutput", string(data)}, " ")
}
