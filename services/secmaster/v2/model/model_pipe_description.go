package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PipeDescription 管道描述
type PipeDescription struct {
}

func (o PipeDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipeDescription struct{}"
	}

	return strings.Join([]string{"PipeDescription", string(data)}, " ")
}
