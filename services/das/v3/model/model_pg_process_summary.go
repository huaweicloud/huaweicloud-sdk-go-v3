package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PgProcessSummary PgProcessSummary
type PgProcessSummary struct {

	// 参数名
	Key *string `json:"key,omitempty"`

	// 参数值
	Value *int64 `json:"value,omitempty"`
}

func (o PgProcessSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PgProcessSummary struct{}"
	}

	return strings.Join([]string{"PgProcessSummary", string(data)}, " ")
}
