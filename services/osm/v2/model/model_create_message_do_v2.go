package model

import (
	"encoding/json"

	"strings"
)

type CreateMessageDoV2 struct {
	// 留言内容

	Content string `json:"content"`
	// 是否授权

	IsAuthorized *int32 `json:"is_authorized,omitempty"`
	// 机密信息

	AuthorizationContent *string `json:"authorization_content,omitempty"`
	// 附件id

	AccessoryIds *[]string `json:"accessory_ids,omitempty"`
}

func (o CreateMessageDoV2) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMessageDoV2 struct{}"
	}

	return strings.Join([]string{"CreateMessageDoV2", string(data)}, " ")
}
