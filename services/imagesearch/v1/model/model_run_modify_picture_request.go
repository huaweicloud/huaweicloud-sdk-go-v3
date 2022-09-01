package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunModifyPictureRequest struct {

	// 实例名称。
	InstanceName string `json:"instance_name" xml:"instance_name"`

	Body *RunModifyPictureReq `json:"body,omitempty" xml:"body"`
}

func (o RunModifyPictureRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunModifyPictureRequest struct{}"
	}

	return strings.Join([]string{"RunModifyPictureRequest", string(data)}, " ")
}
