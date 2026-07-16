package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InputDataInfoRespDataset **参数解释**：数据输入信息为数据集。
type InputDataInfoRespDataset struct {

	// **参数解释**：训练作业的数据集ID。 **取值范围**：不涉及。
	Id string `json:"id"`

	// **参数解释**：训练作业的数据集版本ID。 **约束限制**：使用旧版数据集即service_type不为V3时必填。 **取值范围**：不涉及。
	VersionId *string `json:"version_id,omitempty"`

	// **参数解释**：训练作业需要的数据集OBS路径URL，ModelArts会通过数据集ID和数据集版本ID自动解析生成。如：“/usr/data/”。 **取值范围**：不涉及。
	ObsUrl *string `json:"obs_url,omitempty"`

	// **参数解释**：数据集服务类型。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：取值为V3时表示使用的是资产服务提供的数据集，其他表示旧版数据集。
	ServiceType *string `json:"service_type,omitempty"`

	// **参数解释**：训练作业的数据集名称。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`
}

func (o InputDataInfoRespDataset) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InputDataInfoRespDataset struct{}"
	}

	return strings.Join([]string{"InputDataInfoRespDataset", string(data)}, " ")
}
