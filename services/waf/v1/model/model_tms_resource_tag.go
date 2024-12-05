package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TmsResourceTag TMS标签信息详情
type TmsResourceTag struct {

	// 标签键
	Key *string `json:"key,omitempty"`

	// 标签值
	Value *string `json:"value,omitempty"`
}

func (o TmsResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsResourceTag struct{}"
	}

	return strings.Join([]string{"TmsResourceTag", string(data)}, " ")
}
