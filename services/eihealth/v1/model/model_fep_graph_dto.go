package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FepGraphDto struct {

	// 中心配体名称
	CenterId string `json:"center_id"`

	// 配体对列表
	Pairs []SimilarityDto `json:"pairs"`
}

func (o FepGraphDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FepGraphDto struct{}"
	}

	return strings.Join([]string{"FepGraphDto", string(data)}, " ")
}
