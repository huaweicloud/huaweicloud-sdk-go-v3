package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TopPathSummary topPath详情数据
type TopPathSummary struct {

	// path值。
	Path *string `json:"path,omitempty"`

	// 对应查询类型的值。（流量单位：Byte）
	Value *int64 `json:"value,omitempty"`
}

func (o TopPathSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopPathSummary struct{}"
	}

	return strings.Join([]string{"TopPathSummary", string(data)}, " ")
}
