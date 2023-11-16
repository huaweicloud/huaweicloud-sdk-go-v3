package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBridgeVersionsRequest Request Object
type ListBridgeVersionsRequest struct {

	// servicebridge 类型，可选  rds, mqs 或 cache
	Type string `json:"type"`
}

func (o ListBridgeVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBridgeVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListBridgeVersionsRequest", string(data)}, " ")
}
