package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTokenRequest Request Object
type ShowTokenRequest struct {

	// 应用id。
	BusinessId string `json:"business_id"`

	// 应用id。
	XBusinessId int32 `json:"x-business-id"`
}

func (o ShowTokenRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTokenRequest struct{}"
	}

	return strings.Join([]string{"ShowTokenRequest", string(data)}, " ")
}
