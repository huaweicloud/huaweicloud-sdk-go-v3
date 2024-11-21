package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVolumeRequest Request Object
type CreateVolumeRequest struct {

	// 请求的幂等标识。
	XClientToken *string `json:"X-Client-Token,omitempty"`

	Body *CreateVolumeRequestBody `json:"body,omitempty"`
}

func (o CreateVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVolumeRequest struct{}"
	}

	return strings.Join([]string{"CreateVolumeRequest", string(data)}, " ")
}
