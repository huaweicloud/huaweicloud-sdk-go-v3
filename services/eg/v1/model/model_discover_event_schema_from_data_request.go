package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DiscoverEventSchemaFromDataRequest struct {
	Body *DiscoverEventSchemaFromDataReq `json:"body,omitempty"`
}

func (o DiscoverEventSchemaFromDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiscoverEventSchemaFromDataRequest struct{}"
	}

	return strings.Join([]string{"DiscoverEventSchemaFromDataRequest", string(data)}, " ")
}
