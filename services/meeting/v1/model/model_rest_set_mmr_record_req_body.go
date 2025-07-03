package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestSetMmrRecordReqBody rest设置MMR录播启停请求体
type RestSetMmrRecordReqBody struct {

	// 0：暂停MMR会议录制 1：启动MMR会议录制 2：停止MMR会议录制
	RecordType *int32 `json:"recordType,omitempty"`
}

func (o RestSetMmrRecordReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestSetMmrRecordReqBody struct{}"
	}

	return strings.Join([]string{"RestSetMmrRecordReqBody", string(data)}, " ")
}
