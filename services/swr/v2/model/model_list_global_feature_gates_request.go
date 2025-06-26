package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGlobalFeatureGatesRequest Request Object
type ListGlobalFeatureGatesRequest struct {
}

func (o ListGlobalFeatureGatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalFeatureGatesRequest struct{}"
	}

	return strings.Join([]string{"ListGlobalFeatureGatesRequest", string(data)}, " ")
}
