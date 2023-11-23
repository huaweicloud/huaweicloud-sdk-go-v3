package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFunctionAppRequest Request Object
type ShowFunctionAppRequest struct {

	// 应用ID。
	Id string `json:"id"`
}

func (o ShowFunctionAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFunctionAppRequest struct{}"
	}

	return strings.Join([]string{"ShowFunctionAppRequest", string(data)}, " ")
}
