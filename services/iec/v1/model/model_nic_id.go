package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NicId 网卡ID信息。
type NicId struct {

	// 网卡ID。
	Id string `json:"id"`
}

func (o NicId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NicId struct{}"
	}

	return strings.Join([]string{"NicId", string(data)}, " ")
}
