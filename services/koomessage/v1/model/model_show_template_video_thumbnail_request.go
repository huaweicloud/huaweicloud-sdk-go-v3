package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTemplateVideoThumbnailRequest Request Object
type ShowTemplateVideoThumbnailRequest struct {

	// 目标资源ID。
	AimResourceId string `json:"aim_resource_id"`

	// 请求体参数类型，该字段必须设置为：application/json。
	ContentType string `json:"Content-Type"`
}

func (o ShowTemplateVideoThumbnailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplateVideoThumbnailRequest struct{}"
	}

	return strings.Join([]string{"ShowTemplateVideoThumbnailRequest", string(data)}, " ")
}
