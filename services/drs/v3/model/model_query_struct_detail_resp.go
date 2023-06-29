package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryStructDetailResp 灾备初始化对象详情
type QueryStructDetailResp struct {

	// 任务ID
	JobId string `json:"job_id"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMessage *string `json:"error_message,omitempty"`

	StructDetail *QueryFlowCompareDataResp `json:"struct_detail,omitempty"`
}

func (o QueryStructDetailResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryStructDetailResp struct{}"
	}

	return strings.Join([]string{"QueryStructDetailResp", string(data)}, " ")
}
