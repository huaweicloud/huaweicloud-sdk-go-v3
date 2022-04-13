package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowApplicationDetailRequest struct {
	// 应用ID。

	ApplicationId string `json:"application_id"`
}

func (o ShowApplicationDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApplicationDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowApplicationDetailRequest", string(data)}, " ")
}
