package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigFeaturesResponse Response Object
type ListConfigFeaturesResponse struct {

	// **参数解释**： 特性列表。
	Features *[]ListConfigFeatures `json:"features,omitempty"`

	// **参数解释**： 总特性数量。  **取值范围**： 不涉及。
	TotalRecord    *int32 `json:"totalRecord,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListConfigFeaturesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigFeaturesResponse struct{}"
	}

	return strings.Join([]string{"ListConfigFeaturesResponse", string(data)}, " ")
}
