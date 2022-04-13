package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新素材信息请求（只能修改素材名称）
type UpdateMaterialRequestDto struct {
	// 素材名称

	MaterialName *string `json:"materialName,omitempty"`
}

func (o UpdateMaterialRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMaterialRequestDto struct{}"
	}

	return strings.Join([]string{"UpdateMaterialRequestDto", string(data)}, " ")
}
