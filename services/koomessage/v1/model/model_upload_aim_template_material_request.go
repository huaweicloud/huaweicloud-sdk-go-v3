package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UploadAimTemplateMaterialRequest struct {

	// 请求体参数类型，该字段必须设置为：application/json。
	ContentType string `json:"Content-Type"`

	Body *UploadAimTemplateMaterialRequestBody `json:"body,omitempty"`
}

func (o UploadAimTemplateMaterialRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadAimTemplateMaterialRequest struct{}"
	}

	return strings.Join([]string{"UploadAimTemplateMaterialRequest", string(data)}, " ")
}
