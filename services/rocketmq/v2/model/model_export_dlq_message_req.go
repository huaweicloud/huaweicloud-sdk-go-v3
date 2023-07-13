package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExportDlqMessageReq struct {

	// 主题名称。
	Topic *string `json:"topic,omitempty"`

	// 消息ID列表。
	MsgIdList *[]string `json:"msg_id_list,omitempty"`

	// 唯一Key列表。
	UniqKeyList *[]string `json:"uniq_key_list,omitempty"`
}

func (o ExportDlqMessageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportDlqMessageReq struct{}"
	}

	return strings.Join([]string{"ExportDlqMessageReq", string(data)}, " ")
}
