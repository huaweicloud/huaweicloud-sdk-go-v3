package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ArrowSchema 定义了Arrow Schema的结构，遵循Apache Arrow标准，包含字段定义和元数据信息。
type ArrowSchema struct {

	// Arrow字段列表，定义表的所有列及其类型信息。
	Fields []ArrowField `json:"fields"`

	// Schema的元数据信息，key-value形式的附加信息。
	Metadata map[string]string `json:"metadata,omitempty"`
}

func (o ArrowSchema) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ArrowSchema struct{}"
	}

	return strings.Join([]string{"ArrowSchema", string(data)}, " ")
}
