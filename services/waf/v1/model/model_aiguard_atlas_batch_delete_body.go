package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AiguardAtlasBatchDeleteBody struct {

	// 图库ID列表
	AtlasIds []string `json:"atlas_ids"`
}

func (o AiguardAtlasBatchDeleteBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiguardAtlasBatchDeleteBody struct{}"
	}

	return strings.Join([]string{"AiguardAtlasBatchDeleteBody", string(data)}, " ")
}
