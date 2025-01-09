package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppDetailRequest Request Object
type ShowAppDetailRequest struct {

	// 应用组ID。
	AppGroupId string `json:"app_group_id"`

	// 应用ID。
	AppId string `json:"app_id"`
}

func (o ShowAppDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowAppDetailRequest", string(data)}, " ")
}
