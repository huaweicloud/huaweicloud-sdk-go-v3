package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CountEipAvailableResourcesRequest struct {
	Body *EipResourcesAvailableV3RequestBody `json:"body,omitempty"`
}

func (o CountEipAvailableResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountEipAvailableResourcesRequest struct{}"
	}

	return strings.Join([]string{"CountEipAvailableResourcesRequest", string(data)}, " ")
}
