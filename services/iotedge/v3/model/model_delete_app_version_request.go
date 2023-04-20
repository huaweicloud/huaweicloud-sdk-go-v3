package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteAppVersionRequest struct {

	// 应用ID
	AppId string `json:"app_id"`

	// 应用版本
	Version string `json:"version"`
}

func (o DeleteAppVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppVersionRequest struct{}"
	}

	return strings.Join([]string{"DeleteAppVersionRequest", string(data)}, " ")
}
