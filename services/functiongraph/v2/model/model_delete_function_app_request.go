package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFunctionAppRequest Request Object
type DeleteFunctionAppRequest struct {

	// 应用ID。
	Id string `json:"id"`
}

func (o DeleteFunctionAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFunctionAppRequest struct{}"
	}

	return strings.Join([]string{"DeleteFunctionAppRequest", string(data)}, " ")
}
