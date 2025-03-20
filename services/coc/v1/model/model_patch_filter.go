package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PatchFilter 补丁批准规则过滤器
type PatchFilter struct {

	// key
	Key string `json:"key"`

	// 值
	Values []string `json:"values"`
}

func (o PatchFilter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PatchFilter struct{}"
	}

	return strings.Join([]string{"PatchFilter", string(data)}, " ")
}
