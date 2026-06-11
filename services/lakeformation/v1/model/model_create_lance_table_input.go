package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLanceTableInput Lance表格式结构体
type CreateLanceTableInput struct {
	Schema *LanceSchema `json:"schema,omitempty"`
}

func (o CreateLanceTableInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLanceTableInput struct{}"
	}

	return strings.Join([]string{"CreateLanceTableInput", string(data)}, " ")
}
