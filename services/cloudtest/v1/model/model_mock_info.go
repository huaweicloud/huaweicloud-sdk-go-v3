package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MockInfo struct {

	// mock id
	MockId *string `json:"mock_id,omitempty"`

	// mock开关
	MockSwitch *bool `json:"mock_switch,omitempty"`

	// mock url
	MockUrl *string `json:"mock_url,omitempty"`
}

func (o MockInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MockInfo struct{}"
	}

	return strings.Join([]string{"MockInfo", string(data)}, " ")
}
