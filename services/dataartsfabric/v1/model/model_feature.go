package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Feature 用户特性。
type Feature struct {

	// 配置项key。
	Key *string `json:"key,omitempty"`

	// 配置项value。
	Value *string `json:"value,omitempty"`
}

func (o Feature) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Feature struct{}"
	}

	return strings.Join([]string{"Feature", string(data)}, " ")
}
