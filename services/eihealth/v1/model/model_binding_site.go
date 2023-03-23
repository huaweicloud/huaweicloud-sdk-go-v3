package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 结合位点
type BindingSite struct {

	// 蛋白质3D结构，使用gzip压缩然后转base64格式
	Protein *string `json:"protein,omitempty"`

	BoundingBox *BoundingBox `json:"bounding_box,omitempty"`
}

func (o BindingSite) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindingSite struct{}"
	}

	return strings.Join([]string{"BindingSite", string(data)}, " ")
}
