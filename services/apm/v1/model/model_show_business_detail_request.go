package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBusinessDetailRequest Request Object
type ShowBusinessDetailRequest struct {

	// 应用id。
	BusinessId int64 `json:"business_id"`

	// 应用id。
	XBusinessId int64 `json:"x-business-id"`
}

func (o ShowBusinessDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBusinessDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowBusinessDetailRequest", string(data)}, " ")
}
