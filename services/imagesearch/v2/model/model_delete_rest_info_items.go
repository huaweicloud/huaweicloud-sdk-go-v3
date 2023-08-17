package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteRestInfoItems struct {

	// 数据唯一ID。
	Id *string `json:"id,omitempty"`
}

func (o DeleteRestInfoItems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRestInfoItems struct{}"
	}

	return strings.Join([]string{"DeleteRestInfoItems", string(data)}, " ")
}
