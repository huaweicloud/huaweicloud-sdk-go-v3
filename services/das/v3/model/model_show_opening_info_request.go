package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOpeningInfoRequest Request Object
type ShowOpeningInfoRequest struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o ShowOpeningInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOpeningInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowOpeningInfoRequest", string(data)}, " ")
}
