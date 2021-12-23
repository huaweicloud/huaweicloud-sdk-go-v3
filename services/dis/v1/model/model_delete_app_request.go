package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteAppRequest struct {
	// 需要删除的App名称。

	AppName string `json:"app_name"`
}

func (o DeleteAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppRequest struct{}"
	}

	return strings.Join([]string{"DeleteAppRequest", string(data)}, " ")
}
