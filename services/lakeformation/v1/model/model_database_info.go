package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// database input when grant policy
type DatabaseInfo struct {

	// 数据库名
	Name string `json:"name"`

	// 子表信息
	Tables *[]TableInfo `json:"tables,omitempty"`

	// 子方法信息
	Functions *[]FunctionInfo `json:"functions,omitempty"`
}

func (o DatabaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseInfo struct{}"
	}

	return strings.Join([]string{"DatabaseInfo", string(data)}, " ")
}
