package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpgradeGraph2Request struct {

	// 图ID。
	GraphId string `json:"graph_id"`

	Body *UpgradeGraphReq `json:"body,omitempty"`
}

func (o UpgradeGraph2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeGraph2Request struct{}"
	}

	return strings.Join([]string{"UpgradeGraph2Request", string(data)}, " ")
}
