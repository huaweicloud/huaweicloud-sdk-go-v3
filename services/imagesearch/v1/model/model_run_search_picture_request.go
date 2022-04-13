package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunSearchPictureRequest struct {
	// 实例名称。

	InstanceName string `json:"instance_name"`

	Body *SearchPictureReq `json:"body,omitempty"`
}

func (o RunSearchPictureRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSearchPictureRequest struct{}"
	}

	return strings.Join([]string{"RunSearchPictureRequest", string(data)}, " ")
}
