package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAccessPreviewReqBody struct {
	Configurations *Configuration `json:"configurations"`

	// 资源的唯一资源标识符。
	ResourceUrn string `json:"resource_urn"`
}

func (o CreateAccessPreviewReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccessPreviewReqBody struct{}"
	}

	return strings.Join([]string{"CreateAccessPreviewReqBody", string(data)}, " ")
}
