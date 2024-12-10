package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowShareFeatureGatesRequest Request Object
type ShowShareFeatureGatesRequest struct {
}

func (o ShowShareFeatureGatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowShareFeatureGatesRequest struct{}"
	}

	return strings.Join([]string{"ShowShareFeatureGatesRequest", string(data)}, " ")
}
