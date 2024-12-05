package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HotKeysInfoResponseBody struct {

	// 热Key名。
	Name *string `json:"name,omitempty"`

	// 热Key类型。
	Type *string `json:"type,omitempty"`

	// 热Key命令。
	Command *string `json:"command,omitempty"`

	// 热Key QPS。
	Qps *int32 `json:"qps,omitempty"`

	// 热key所在的DB。
	DbId *int32 `json:"db_id,omitempty"`
}

func (o HotKeysInfoResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HotKeysInfoResponseBody struct{}"
	}

	return strings.Join([]string{"HotKeysInfoResponseBody", string(data)}, " ")
}
