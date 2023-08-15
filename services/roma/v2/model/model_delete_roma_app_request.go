package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRomaAppRequest Request Object
type DeleteRomaAppRequest struct {

	// 应用ID
	AppId string `json:"app_id"`

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o DeleteRomaAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRomaAppRequest struct{}"
	}

	return strings.Join([]string{"DeleteRomaAppRequest", string(data)}, " ")
}
