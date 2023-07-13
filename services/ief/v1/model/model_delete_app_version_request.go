package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAppVersionRequest Request Object
type DeleteAppVersionRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 应用ID
	AppId string `json:"app_id"`

	// 版本ID
	VersionId string `json:"version_id"`
}

func (o DeleteAppVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppVersionRequest struct{}"
	}

	return strings.Join([]string{"DeleteAppVersionRequest", string(data)}, " ")
}
