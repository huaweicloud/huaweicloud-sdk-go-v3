package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowSubBusinessDetailRequest struct {

	// 子应用id。
	SubBusinessId int64 `json:"sub_business_id"`

	// 应用id。
	XBusinessId int64 `json:"x-business-id"`
}

func (o ShowSubBusinessDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubBusinessDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowSubBusinessDetailRequest", string(data)}, " ")
}
