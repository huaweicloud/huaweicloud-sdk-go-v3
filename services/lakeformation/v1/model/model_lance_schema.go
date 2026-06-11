package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LanceSchema 定义了Arrow Schema的结构，遵循Apache Arrow标准，包含字段定义和元数据信息。
type LanceSchema struct {

	// Arrow字段列表，定义表的所有列及其类型信息。
	Fields []ArrowField `json:"fields"`

	// Schema的元数据信息，key-value形式的附加信息。
	Metadata map[string]string `json:"metadata,omitempty"`
}

func (o LanceSchema) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LanceSchema struct{}"
	}

	return strings.Join([]string{"LanceSchema", string(data)}, " ")
}
