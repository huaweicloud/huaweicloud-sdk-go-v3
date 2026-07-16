package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAlgorithmVersionToGalleryBody 创建发布算法资产请求
type CreateAlgorithmVersionToGalleryBody struct {

	// **参数解释**：资产id。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ContentId *string `json:"content_id,omitempty"`

	ContentInfo *ContentInfo `json:"content_info,omitempty"`

	Algorithm *AlgorithmInfo `json:"algorithm,omitempty"`
}

func (o CreateAlgorithmVersionToGalleryBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlgorithmVersionToGalleryBody struct{}"
	}

	return strings.Join([]string{"CreateAlgorithmVersionToGalleryBody", string(data)}, " ")
}
