package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Properties **参数解释**：数据集资产扩展信息。 **约束限制**：不涉及。
type Properties struct {

	// **参数解释**：数据集的条数。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	DatasetNum *int32 `json:"dataset_num,omitempty"`

	// **参数解释**：数据集资产资源文件总大小。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	DatasetSize *int64 `json:"dataset_size,omitempty"`

	// **参数解释**：数据集资产评测模版相对地址。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PromptConfig *string `json:"prompt_config,omitempty"`

	// **参数解释**：数据集资产版权信息，包含数据集信息，License和地址。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	DataCopyright *string `json:"data_copyright,omitempty"`

	// **参数解释**：数据集资产索引文件地址。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ManifestFilePath *string `json:"manifest_file_path,omitempty"`

	// **参数解释**：算子参数列表。 **约束限制**：不涉及。
	Arguments *[]OperatorArgument `json:"arguments,omitempty"`

	// **参数解释**：数据集资产修改人信息。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Updater *string `json:"updater,omitempty"`
}

func (o Properties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Properties struct{}"
	}

	return strings.Join([]string{"Properties", string(data)}, " ")
}
