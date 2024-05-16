package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlameLineTreeRequest Request Object
type ShowFlameLineTreeRequest struct {
	Body *FlameLineTreeInfo `json:"body,omitempty"`
}

func (o ShowFlameLineTreeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlameLineTreeRequest struct{}"
	}

	return strings.Join([]string{"ShowFlameLineTreeRequest", string(data)}, " ")
}
