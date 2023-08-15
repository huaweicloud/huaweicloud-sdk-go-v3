package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteNoticeReq struct {

	// 批量删除通知消息id列表
	NoticeIds []string `json:"notice_ids"`
}

func (o BatchDeleteNoticeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteNoticeReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteNoticeReq", string(data)}, " ")
}
