package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateCloudPhoneServerRequest struct {
	Body *CreateCloudPhoneServerRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateCloudPhoneServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudPhoneServerRequest struct{}"
	}

	return strings.Join([]string{"CreateCloudPhoneServerRequest", string(data)}, " ")
}
