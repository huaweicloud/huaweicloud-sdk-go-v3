package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlgorithmResponseJobConfigRemoteConstraints struct {

	// 数据输入类型，包括数据存储位置、数据集两种方式。
	DataType *string `json:"data_type,omitempty"`

	// 数据输入为数据集时的相关属性。枚举值：   - data_format数据格式。   - data_segmentation数据切分方式。   - dataset_type标注类型。
	Attributes *[]map[string]string `json:"attributes,omitempty"`
}

func (o AlgorithmResponseJobConfigRemoteConstraints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlgorithmResponseJobConfigRemoteConstraints struct{}"
	}

	return strings.Join([]string{"AlgorithmResponseJobConfigRemoteConstraints", string(data)}, " ")
}
