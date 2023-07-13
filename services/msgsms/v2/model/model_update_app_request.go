package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAppRequest Request Object
type UpdateAppRequest struct {

	// 应用主键ID
	Id string `json:"id"`

	Body *SmsAppUpdateReq `json:"body,omitempty"`
}

func (o UpdateAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppRequest struct{}"
	}

	return strings.Join([]string{"UpdateAppRequest", string(data)}, " ")
}
