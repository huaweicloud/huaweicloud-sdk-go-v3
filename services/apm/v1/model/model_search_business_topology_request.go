package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchBusinessTopologyRequest Request Object
type SearchBusinessTopologyRequest struct {

	// 应用id。
	XBusinessId int64 `json:"x-business-id"`

	Body *BusinessTopoRequest `json:"body,omitempty"`
}

func (o SearchBusinessTopologyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchBusinessTopologyRequest struct{}"
	}

	return strings.Join([]string{"SearchBusinessTopologyRequest", string(data)}, " ")
}
