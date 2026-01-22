package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigFeaturesRequest Request Object
type ListConfigFeaturesRequest struct {
}

func (o ListConfigFeaturesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigFeaturesRequest struct{}"
	}

	return strings.Join([]string{"ListConfigFeaturesRequest", string(data)}, " ")
}
