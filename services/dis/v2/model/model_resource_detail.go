package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceDetail struct {

	// 标签列表，没有标签默认为空数组
	Action *string `json:"action,omitempty"`
}

func (o ResourceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceDetail struct{}"
	}

	return strings.Join([]string{"ResourceDetail", string(data)}, " ")
}
