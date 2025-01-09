package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAppIconRequest Request Object
type DeleteAppIconRequest struct {

	// 应用组ID。
	AppGroupId string `json:"app_group_id"`

	// 应用ID。
	AppId string `json:"app_id"`
}

func (o DeleteAppIconRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppIconRequest struct{}"
	}

	return strings.Join([]string{"DeleteAppIconRequest", string(data)}, " ")
}
