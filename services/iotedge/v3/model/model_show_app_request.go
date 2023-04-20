package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowAppRequest struct {

	// 应用模板ID。
	AppId string `json:"app_id"`
}

func (o ShowAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppRequest struct{}"
	}

	return strings.Join([]string{"ShowAppRequest", string(data)}, " ")
}
