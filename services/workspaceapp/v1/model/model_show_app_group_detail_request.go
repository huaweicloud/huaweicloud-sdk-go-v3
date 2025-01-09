package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppGroupDetailRequest Request Object
type ShowAppGroupDetailRequest struct {

	// 应用组ID。
	AppGroupId string `json:"app_group_id"`
}

func (o ShowAppGroupDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppGroupDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowAppGroupDetailRequest", string(data)}, " ")
}
