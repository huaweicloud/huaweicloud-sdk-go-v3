package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFeatureChildrenRequest Request Object
type ShowFeatureChildrenRequest struct {
	FeatureId string `json:"feature_id"`

	Body *QueryTestItemTreeInfo `json:"body,omitempty"`
}

func (o ShowFeatureChildrenRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFeatureChildrenRequest struct{}"
	}

	return strings.Join([]string{"ShowFeatureChildrenRequest", string(data)}, " ")
}
