package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资源对象id列表。
type ResourceList struct {

	// 资源ID
	Id string `json:"id"`
}

func (o ResourceList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceList struct{}"
	}

	return strings.Join([]string{"ResourceList", string(data)}, " ")
}
