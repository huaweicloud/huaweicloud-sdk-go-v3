package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 上传模板素材响应体。
type UploadAimTemplateMaterialResponseMode struct {

	// 模板素材ID。
	MaterialId *string `json:"material_id,omitempty"`

	// 资源ID。
	AimResourceId *string `json:"aim_resource_id,omitempty"`
}

func (o UploadAimTemplateMaterialResponseMode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadAimTemplateMaterialResponseMode struct{}"
	}

	return strings.Join([]string{"UploadAimTemplateMaterialResponseMode", string(data)}, " ")
}
