package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyVolumeQoSRequestBody This is a auto create Body Object
type ModifyVolumeQoSRequestBody struct {
	QosModify *ModifyVolumeQoSOption `json:"qos_modify"`
}

func (o ModifyVolumeQoSRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyVolumeQoSRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyVolumeQoSRequestBody", string(data)}, " ")
}
