package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeMasterStandbyAsyncRequest Request Object
type ChangeMasterStandbyAsyncRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ChangeMasterStandbyAsyncRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeMasterStandbyAsyncRequest struct{}"
	}

	return strings.Join([]string{"ChangeMasterStandbyAsyncRequest", string(data)}, " ")
}
