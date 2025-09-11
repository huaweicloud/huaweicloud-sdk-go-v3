package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceTagInfoBean struct {

	// 标签KEY
	Key *string `json:"key,omitempty"`

	// 标签VALUE
	Value string `json:"value"`
}

func (o ResourceTagInfoBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTagInfoBean struct{}"
	}

	return strings.Join([]string{"ResourceTagInfoBean", string(data)}, " ")
}
