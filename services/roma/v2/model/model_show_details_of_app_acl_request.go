package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDetailsOfAppAclRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 应用编号

	AppId string `json:"app_id"`
}

func (o ShowDetailsOfAppAclRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfAppAclRequest struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfAppAclRequest", string(data)}, " ")
}
