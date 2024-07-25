package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionUpdateViewDtoMultiViewModelViewUpdateAttrDto struct {

	// **参数解释：**  版本对象ID。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	VersionId string `json:"versionId"`

	NewViewAttrs *MultiViewModelViewUpdateAttrDto `json:"newViewAttrs"`

	// **参数解释：**  指定不复制的视图属性，其值将被设置为null。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	NeedSetNull *[]string `json:"needSetNull,omitempty"`
}

func (o VersionUpdateViewDtoMultiViewModelViewUpdateAttrDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionUpdateViewDtoMultiViewModelViewUpdateAttrDto struct{}"
	}

	return strings.Join([]string{"VersionUpdateViewDtoMultiViewModelViewUpdateAttrDto", string(data)}, " ")
}
