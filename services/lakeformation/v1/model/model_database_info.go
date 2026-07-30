package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DatabaseInfo database input when grant policy
type DatabaseInfo struct {

	// 数据库名
	Name string `json:"name"`

	// 子表信息
	Tables *[]TableInfo `json:"tables,omitempty"`

	// 子方法信息
	Functions *[]FunctionInfo `json:"functions,omitempty"`

	// **参数解释:** 子数据集信息 **约束限制:** 数组元素个数为0~10 **取值范围:** 数组元素个数为0~10 **默认取值:** 不涉及
	Datasets *[]DatasetInfo `json:"datasets,omitempty"`
}

func (o DatabaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseInfo struct{}"
	}

	return strings.Join([]string{"DatabaseInfo", string(data)}, " ")
}
