package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CatalogContext **参数解释**：LakeFormation相关信息。 **约束限制**：不涉及。
type CatalogContext struct {

	// **参数解释**：LakeFormation的catalogName。 **约束限制**：不涉及。 **取值范围**：长度为1~255个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	CatalogName string `json:"catalog_name"`
}

func (o CatalogContext) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CatalogContext struct{}"
	}

	return strings.Join([]string{"CatalogContext", string(data)}, " ")
}
