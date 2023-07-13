package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteTemplateMaterialRequestBody struct {

	// 模板素材ID数组。
	MaterialIds []string `json:"material_ids"`
}

func (o DeleteTemplateMaterialRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTemplateMaterialRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteTemplateMaterialRequestBody", string(data)}, " ")
}
