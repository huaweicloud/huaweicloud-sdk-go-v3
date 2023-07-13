package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Id id
type Id struct {

	// 关联监听器ID。
	Id string `json:"id"`
}

func (o Id) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Id struct{}"
	}

	return strings.Join([]string{"Id", string(data)}, " ")
}
