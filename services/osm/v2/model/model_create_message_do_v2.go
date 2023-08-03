package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateMessageDoV2 struct {

	// 留言内容
	Content string `json:"content"`

	// 附件id
	AccessoryIds *[]string `json:"accessory_ids,omitempty"`
}

func (o CreateMessageDoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMessageDoV2 struct{}"
	}

	return strings.Join([]string{"CreateMessageDoV2", string(data)}, " ")
}
