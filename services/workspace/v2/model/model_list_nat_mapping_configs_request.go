package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNatMappingConfigsRequest Request Object
type ListNatMappingConfigsRequest struct {

	// 站点ID
	SiteId *string `json:"site_id,omitempty"`
}

func (o ListNatMappingConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNatMappingConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListNatMappingConfigsRequest", string(data)}, " ")
}
