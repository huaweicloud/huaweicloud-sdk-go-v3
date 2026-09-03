package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DatasetConfig struct {

	// 训练数据集名称，取自数据集列表接口响应体name。
	DatasetName *string `json:"dataset_name,omitempty"`

	// 所使用的数据集来源，取值datamng|OBS|DB,分别表示来自于数据工程|OBS|数据库
	DatasetSource *string `json:"dataset_source,omitempty"`

	// 训练数据集id，取自数据集列表接口响应体dataset_id。
	DatasetId *string `json:"dataset_id,omitempty"`

	// 训练、验证数据集分割比率，当该模型支持验证集且验证集来自选择的训练集时使用，取值大于等于1，小于等于50。
	SplitRatio *int32 `json:"split_ratio,omitempty"`

	// 数据集使用的阶段，取值为train|eval|test，分别表示该数据集用于训练|验证|测试。
	UsedStep *string `json:"used_step,omitempty"`

	// 数据集配比比率，表示使用多少比率的该数据集进行训练。
	DatasetProportion *int32 `json:"dataset_proportion,omitempty"`
}

func (o DatasetConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatasetConfig struct{}"
	}

	return strings.Join([]string{"DatasetConfig", string(data)}, " ")
}
