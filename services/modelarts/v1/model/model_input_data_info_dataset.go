package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InputDataInfoDataset 数据输入信息为数据集。
type InputDataInfoDataset struct {

	// 训练作业的数据集ID。
	Id string `json:"id"`

	// 训练作业的数据集版本ID。
	VersionId *string `json:"version_id,omitempty"`

	// 训练作业需要的数据集OBS路径URL，ModelArts会通过数据集ID和数据集版本ID自动解析生成。如：“/usr/data/”。
	ObsUrl *string `json:"obs_url,omitempty"`

	// **参数解释**：数据集服务类型。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：取值为V3时表示使用的是资产服务提供的数据集，其他表示旧版数据集。
	ServiceType *string `json:"service_type,omitempty"`

	// **参数解释**：训练作业的数据集名称。
	Name *string `json:"name,omitempty"`

	// **参数解释**：精调训练作业的数据集配比比率，表示使用多少比率的该数据集进行训练。
	DatasetProportion *int32 `json:"dataset_proportion,omitempty"`
}

func (o InputDataInfoDataset) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InputDataInfoDataset struct{}"
	}

	return strings.Join([]string{"InputDataInfoDataset", string(data)}, " ")
}
