package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowInstanceStatusInfoRequest struct {

	// CloudIDE实例id
	InstanceId string `json:"instance_id" xml:"instance_id"`
}

func (o ShowInstanceStatusInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceStatusInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceStatusInfoRequest", string(data)}, " ")
}
