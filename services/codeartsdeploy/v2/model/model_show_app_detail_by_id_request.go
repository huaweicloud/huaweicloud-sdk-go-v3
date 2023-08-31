package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppDetailByIdRequest Request Object
type ShowAppDetailByIdRequest struct {

	// 应用id
	AppId string `json:"app_id"`
}

func (o ShowAppDetailByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppDetailByIdRequest struct{}"
	}

	return strings.Join([]string{"ShowAppDetailByIdRequest", string(data)}, " ")
}
