package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlterDatasetInput 用户输入的数据集
type AlterDatasetInput struct {

	// 数据集的描述信息
	Description *string `json:"description,omitempty"`

	DatasetFormat *DatasetFileFormat `json:"dataset_format,omitempty"`

	// 数据集其他属性
	Properties map[string]string `json:"properties,omitempty"`
}

func (o AlterDatasetInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlterDatasetInput struct{}"
	}

	return strings.Join([]string{"AlterDatasetInput", string(data)}, " ")
}
