package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunQueryDocumentModerationJobRequest Request Object
type RunQueryDocumentModerationJobRequest struct {

	// 创建作业成功时，接口返回的job_id。
	JobId string `json:"job_id"`
}

func (o RunQueryDocumentModerationJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunQueryDocumentModerationJobRequest struct{}"
	}

	return strings.Join([]string{"RunQueryDocumentModerationJobRequest", string(data)}, " ")
}
