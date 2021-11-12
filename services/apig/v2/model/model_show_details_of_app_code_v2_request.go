package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDetailsOfAppCodeV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 应用编号

	AppId string `json:"app_id"`
	// APP Code编号

	AppCodeId string `json:"app_code_id"`
}

func (o ShowDetailsOfAppCodeV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfAppCodeV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfAppCodeV2Request", string(data)}, " ")
}
