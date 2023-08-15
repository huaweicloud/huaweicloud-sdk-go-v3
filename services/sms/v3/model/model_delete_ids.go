package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIds 批量删除参数
type DeleteIds struct {

	// 所有删除对象ID的集合
	Ids []string `json:"ids"`
}

func (o DeleteIds) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIds struct{}"
	}

	return strings.Join([]string{"DeleteIds", string(data)}, " ")
}
