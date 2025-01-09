package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstallAppRequest Request Object
type InstallAppRequest struct {

	// 应用ID。
	AppId string `json:"app_id"`

	Body *AutoInstallAppReq `json:"body,omitempty"`
}

func (o InstallAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstallAppRequest struct{}"
	}

	return strings.Join([]string{"InstallAppRequest", string(data)}, " ")
}
