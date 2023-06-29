package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAppCodeV2Request Request Object
type DeleteAppCodeV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 应用编号
	AppId string `json:"app_id"`

	// APP Code编号
	AppCodeId string `json:"app_code_id"`
}

func (o DeleteAppCodeV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppCodeV2Request struct{}"
	}

	return strings.Join([]string{"DeleteAppCodeV2Request", string(data)}, " ")
}
