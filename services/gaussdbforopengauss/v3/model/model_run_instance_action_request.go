package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunInstanceActionRequest Request Object
type RunInstanceActionRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *OpenGaussInstanceActionRequest `json:"body,omitempty"`
}

func (o RunInstanceActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunInstanceActionRequest struct{}"
	}

	return strings.Join([]string{"RunInstanceActionRequest", string(data)}, " ")
}
