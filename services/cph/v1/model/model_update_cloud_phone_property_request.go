package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateCloudPhonePropertyRequest struct {
	Body *Phones `json:"body,omitempty"`
}

func (o UpdateCloudPhonePropertyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCloudPhonePropertyRequest struct{}"
	}

	return strings.Join([]string{"UpdateCloudPhonePropertyRequest", string(data)}, " ")
}
