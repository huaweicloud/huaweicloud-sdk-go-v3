package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MemberRef
type MemberRef struct {

	// 参数解释：后端服务器ID。
	Id string `json:"id"`
}

func (o MemberRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberRef struct{}"
	}

	return strings.Join([]string{"MemberRef", string(data)}, " ")
}
