package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddFeatureRequest Request Object
type AddFeatureRequest struct {
	Body *AddTestItemInfo `json:"body,omitempty"`
}

func (o AddFeatureRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddFeatureRequest struct{}"
	}

	return strings.Join([]string{"AddFeatureRequest", string(data)}, " ")
}
