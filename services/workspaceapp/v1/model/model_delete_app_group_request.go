package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAppGroupRequest Request Object
type DeleteAppGroupRequest struct {

	// 应用组ID。
	AppGroupId string `json:"app_group_id"`
}

func (o DeleteAppGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteAppGroupRequest", string(data)}, " ")
}
