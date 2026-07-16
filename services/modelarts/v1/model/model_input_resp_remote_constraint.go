package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InputRespRemoteConstraint struct {

	// **参数解释**：数据输入类型，包括数据存储位置、数据集两种方式。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	DataType *string `json:"data_type,omitempty"`

	// **参数解释**：相关属性。 **约束限制**：不涉及。 **取值范围**： 数据输入为数据集时：   - data_format：数据格式   - data_segmentation：数据切分方式   - dataset_type：标注类型  **默认取值**：不涉及。
	Attributes *string `json:"attributes,omitempty"`
}

func (o InputRespRemoteConstraint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InputRespRemoteConstraint struct{}"
	}

	return strings.Join([]string{"InputRespRemoteConstraint", string(data)}, " ")
}
