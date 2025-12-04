package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindEipRequest Request Object
type BindEipRequest struct {

	// 实例 ID。
	InstanceId string `json:"instance_id"`

	Body *BindEipOpenRequest `json:"body,omitempty"`
}

func (o BindEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindEipRequest struct{}"
	}

	return strings.Join([]string{"BindEipRequest", string(data)}, " ")
}
