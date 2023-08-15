package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagCreate struct {

	// tag标签名称。
	Name string `json:"name"`

	// tag标签描述信息。
	Description string `json:"description"`
}

func (o TagCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagCreate struct{}"
	}

	return strings.Join([]string{"TagCreate", string(data)}, " ")
}
