package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CheckRomaInstanceListV2Request struct {

	// 实例状态
	Status *string `json:"status,omitempty" xml:"status"`

	// 偏移量，大于等于0
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty" xml:"limit"`
}

func (o CheckRomaInstanceListV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRomaInstanceListV2Request struct{}"
	}

	return strings.Join([]string{"CheckRomaInstanceListV2Request", string(data)}, " ")
}
