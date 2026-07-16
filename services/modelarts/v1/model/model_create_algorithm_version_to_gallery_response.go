package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAlgorithmVersionToGalleryResponse Response Object
type CreateAlgorithmVersionToGalleryResponse struct {

	// **参数解释**：资产id。 **取值范围**：不涉及。
	ContentId *string `json:"content_id,omitempty"`

	// **参数解释**：版本号id。 **取值范围**：不涉及。
	VersionId *string `json:"version_id,omitempty"`

	// **参数解释**：版本数量。 **取值范围**：不涉及。
	VersionNum *string `json:"version_num,omitempty"`

	// **参数解释**：资产uri地址。 **取值范围**：不涉及。
	ContentUri     *string `json:"content_uri,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAlgorithmVersionToGalleryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlgorithmVersionToGalleryResponse struct{}"
	}

	return strings.Join([]string{"CreateAlgorithmVersionToGalleryResponse", string(data)}, " ")
}
