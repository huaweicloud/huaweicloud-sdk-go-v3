package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFeaturesRequest Request Object
type ListFeaturesRequest struct {

	// **参数解释**：特性名称。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： - NOTEBOOK：用户显式创建的Notebook实例。  **默认取值**：NOTEBOOK。
	Feature string `json:"feature"`
}

func (o ListFeaturesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFeaturesRequest struct{}"
	}

	return strings.Join([]string{"ListFeaturesRequest", string(data)}, " ")
}
