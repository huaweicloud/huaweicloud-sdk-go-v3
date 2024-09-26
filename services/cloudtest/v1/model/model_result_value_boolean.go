package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResultValueBoolean 请求的返回的数据对象
type ResultValueBoolean struct {

	// 实际的数据类型：单个对象，集合 或 NULL
	Value *bool `json:"value,omitempty"`
}

func (o ResultValueBoolean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResultValueBoolean struct{}"
	}

	return strings.Join([]string{"ResultValueBoolean", string(data)}, " ")
}
