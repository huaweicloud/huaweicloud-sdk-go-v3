package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAppRequest Request Object
type DeleteAppRequest struct {

	// 应用ID
	AppId string `json:"app_id"`
}

func (o DeleteAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppRequest struct{}"
	}

	return strings.Join([]string{"DeleteAppRequest", string(data)}, " ")
}
