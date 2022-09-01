package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 灾备初始化对象详情
type QueryStructDetailResp struct {

	// 任务ID
	JobId string `json:"job_id" xml:"job_id"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 错误信息
	ErrorMessage *string `json:"error_message,omitempty" xml:"error_message"`

	StructDetail *QueryFlowCompareDataResp `json:"struct_detail,omitempty" xml:"struct_detail"`
}

func (o QueryStructDetailResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryStructDetailResp struct{}"
	}

	return strings.Join([]string{"QueryStructDetailResp", string(data)}, " ")
}
