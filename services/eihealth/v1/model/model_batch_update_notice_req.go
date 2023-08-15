package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdateNoticeReq struct {

	// 批量更新通知消息id列表
	NoticeIds []string `json:"notice_ids"`

	Operation *NoticeOperation `json:"operation"`
}

func (o BatchUpdateNoticeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateNoticeReq struct{}"
	}

	return strings.Join([]string{"BatchUpdateNoticeReq", string(data)}, " ")
}
