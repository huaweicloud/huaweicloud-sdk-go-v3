package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCloudPhoneSingleServerRequest Request Object
type CreateCloudPhoneSingleServerRequest struct {
	Body *CreateCloudPhoneSingleServerRequestBody `json:"body,omitempty"`
}

func (o CreateCloudPhoneSingleServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudPhoneSingleServerRequest struct{}"
	}

	return strings.Join([]string{"CreateCloudPhoneSingleServerRequest", string(data)}, " ")
}
