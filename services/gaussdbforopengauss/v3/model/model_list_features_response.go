package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFeaturesResponse Response Object
type ListFeaturesResponse struct {

	// 高级特性列表。
	Features       *[]FeatureResult `json:"features,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListFeaturesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFeaturesResponse struct{}"
	}

	return strings.Join([]string{"ListFeaturesResponse", string(data)}, " ")
}
