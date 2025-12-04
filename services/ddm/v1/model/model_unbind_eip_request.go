package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnbindEipRequest Request Object
type UnbindEipRequest struct {

	// 实例 ID。
	InstanceId string `json:"instance_id"`
}

func (o UnbindEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnbindEipRequest struct{}"
	}

	return strings.Join([]string{"UnbindEipRequest", string(data)}, " ")
}
