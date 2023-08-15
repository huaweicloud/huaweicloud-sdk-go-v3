package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobLogRequest struct {

	// 桥接作业id
	BridgeId string `json:"bridge_id"`

	// 分类作业id
	ClassificationId *string `json:"classification_id,omitempty"`
}

func (o JobLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobLogRequest struct{}"
	}

	return strings.Join([]string{"JobLogRequest", string(data)}, " ")
}
