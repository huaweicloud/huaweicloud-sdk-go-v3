package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ChangeMasterStandbyRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ChangeMasterStandbyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeMasterStandbyRequest struct{}"
	}

	return strings.Join([]string{"ChangeMasterStandbyRequest", string(data)}, " ")
}
