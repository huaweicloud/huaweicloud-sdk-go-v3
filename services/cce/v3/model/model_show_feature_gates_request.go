package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFeatureGatesRequest Request Object
type ShowFeatureGatesRequest struct {
}

func (o ShowFeatureGatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFeatureGatesRequest struct{}"
	}

	return strings.Join([]string{"ShowFeatureGatesRequest", string(data)}, " ")
}
