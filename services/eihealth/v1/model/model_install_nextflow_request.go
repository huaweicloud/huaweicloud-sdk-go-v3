package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstallNextflowRequest Request Object
type InstallNextflowRequest struct {
	Body *InstallNextflowRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o InstallNextflowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstallNextflowRequest struct{}"
	}

	return strings.Join([]string{"InstallNextflowRequest", string(data)}, " ")
}
