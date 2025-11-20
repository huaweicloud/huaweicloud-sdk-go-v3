package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Restore2ExistRequest Request Object
type Restore2ExistRequest struct {

	// 实例 ID。
	InstanceId string `json:"instance_id"`

	Body *RestoreInst2ExistReq `json:"body,omitempty"`
}

func (o Restore2ExistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Restore2ExistRequest struct{}"
	}

	return strings.Join([]string{"Restore2ExistRequest", string(data)}, " ")
}
