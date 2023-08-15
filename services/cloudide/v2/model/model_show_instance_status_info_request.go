package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceStatusInfoRequest Request Object
type ShowInstanceStatusInfoRequest struct {

	// CodeArtsIDEOnline实例id
	InstanceId string `json:"instance_id"`
}

func (o ShowInstanceStatusInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceStatusInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceStatusInfoRequest", string(data)}, " ")
}
