package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDetailsOfAppV2Request Request Object
type ShowDetailsOfAppV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 应用编号
	AppId string `json:"app_id"`
}

func (o ShowDetailsOfAppV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfAppV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfAppV2Request", string(data)}, " ")
}
