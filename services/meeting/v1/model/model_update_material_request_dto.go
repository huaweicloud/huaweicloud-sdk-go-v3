package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 信息窗素材信息。
type UpdateMaterialRequestDto struct {

	// 素材名称。
	MaterialName *string `json:"materialName,omitempty"`
}

func (o UpdateMaterialRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMaterialRequestDto struct{}"
	}

	return strings.Join([]string{"UpdateMaterialRequestDto", string(data)}, " ")
}
