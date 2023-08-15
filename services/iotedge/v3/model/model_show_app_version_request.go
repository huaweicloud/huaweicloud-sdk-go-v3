package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppVersionRequest Request Object
type ShowAppVersionRequest struct {

	// 应用ID
	AppId string `json:"app_id"`

	// 应用版本
	Version string `json:"version"`
}

func (o ShowAppVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppVersionRequest struct{}"
	}

	return strings.Join([]string{"ShowAppVersionRequest", string(data)}, " ")
}
