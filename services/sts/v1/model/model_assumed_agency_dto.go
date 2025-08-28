package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssumedAgencyDto 目标委托信息。
type AssumedAgencyDto struct {

	// 目标委托的URN。
	Urn string `json:"urn"`

	// 目标委托的唯一标志，包含了委托ID和委托会话名称信息。
	Id string `json:"id"`
}

func (o AssumedAgencyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssumedAgencyDto struct{}"
	}

	return strings.Join([]string{"AssumedAgencyDto", string(data)}, " ")
}
