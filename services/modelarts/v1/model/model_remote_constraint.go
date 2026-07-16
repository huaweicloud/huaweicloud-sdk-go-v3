package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoteConstraint 数据输入约束。
type RemoteConstraint struct {

	// 数据输入类型，支持数据存储位置（OBS）、ModelArts数据集两种方式。
	DataType *string `json:"data_type,omitempty"`

	// 数据输入为数据集时的相关属性。枚举值：   - data_format：数据格式。   - data_segmentation：数据切分方式。   - dataset_type：标注类型。
	Attributes *[]map[string]string `json:"attributes,omitempty"`
}

func (o RemoteConstraint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoteConstraint struct{}"
	}

	return strings.Join([]string{"RemoteConstraint", string(data)}, " ")
}
