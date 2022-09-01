package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunDeletePictureRequest struct {

	// 实例名称。
	InstanceName string `json:"instance_name" xml:"instance_name"`

	Body *DeletePictureReq `json:"body,omitempty" xml:"body"`
}

func (o RunDeletePictureRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunDeletePictureRequest struct{}"
	}

	return strings.Join([]string{"RunDeletePictureRequest", string(data)}, " ")
}
