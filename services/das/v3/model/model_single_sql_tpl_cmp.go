package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SingleSqlTplCmp SQL模板对比
type SingleSqlTplCmp struct {

	// 趋势指标名称
	Name *string `json:"name,omitempty"`

	// 趋势指标值
	Values *[]float64 `json:"values,omitempty"`
}

func (o SingleSqlTplCmp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SingleSqlTplCmp struct{}"
	}

	return strings.Join([]string{"SingleSqlTplCmp", string(data)}, " ")
}
