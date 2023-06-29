package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryInfo struct {

	// ResourceQL 查询字段
	SelectFields *[]string `json:"select_fields,omitempty"`
}

func (o QueryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryInfo struct{}"
	}

	return strings.Join([]string{"QueryInfo", string(data)}, " ")
}
