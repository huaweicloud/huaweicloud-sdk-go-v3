package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddVrrpConfigRequest Request Object
type AddVrrpConfigRequest struct {

	// 智能企业网关ID
	IegId string `json:"ieg_id"`

	Body *CreateUpdateVrrpConfigRequestBody `json:"body,omitempty"`
}

func (o AddVrrpConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddVrrpConfigRequest struct{}"
	}

	return strings.Join([]string{"AddVrrpConfigRequest", string(data)}, " ")
}
