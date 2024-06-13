package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAllFeatureChildrenRequest Request Object
type ShowAllFeatureChildrenRequest struct {
	FeatureId string `json:"feature_id"`

	Body *QueryTestItemTreeInfo `json:"body,omitempty"`
}

func (o ShowAllFeatureChildrenRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAllFeatureChildrenRequest struct{}"
	}

	return strings.Join([]string{"ShowAllFeatureChildrenRequest", string(data)}, " ")
}
