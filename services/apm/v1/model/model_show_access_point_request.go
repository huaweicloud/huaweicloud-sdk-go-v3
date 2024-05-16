package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAccessPointRequest Request Object
type ShowAccessPointRequest struct {

	// 应用id。
	BusinessId string `json:"business_id"`

	// 应用id。
	XBusinessId int32 `json:"x-business-id"`

	Body *AccessPointModel `json:"body,omitempty"`
}

func (o ShowAccessPointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAccessPointRequest struct{}"
	}

	return strings.Join([]string{"ShowAccessPointRequest", string(data)}, " ")
}
