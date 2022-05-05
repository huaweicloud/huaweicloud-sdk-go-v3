package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Striping struct {

	// 最近领域个数。
	NearestNeighborhood int32 `json:"nearest_neighborhood"`

	// 相似程度。
	Band int32 `json:"band"`

	// 相似距离。
	Row int32 `json:"row"`
}

func (o Striping) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Striping struct{}"
	}

	return strings.Join([]string{"Striping", string(data)}, " ")
}
