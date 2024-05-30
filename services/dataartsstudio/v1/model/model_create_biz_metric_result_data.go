package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBizMetricResultData data，统一的返回结果的最外层数据结构。
type CreateBizMetricResultData struct {
	Value *BizMetricVo `json:"value,omitempty"`
}

func (o CreateBizMetricResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBizMetricResultData struct{}"
	}

	return strings.Join([]string{"CreateBizMetricResultData", string(data)}, " ")
}
