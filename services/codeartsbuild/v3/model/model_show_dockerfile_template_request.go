package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDockerfileTemplateRequest Request Object
type ShowDockerfileTemplateRequest struct {

	// 请求体
	ImageId string `json:"image_id"`
}

func (o ShowDockerfileTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDockerfileTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowDockerfileTemplateRequest", string(data)}, " ")
}
