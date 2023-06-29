package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryRpoAndRtoResp RPO和RTO信息体
type QueryRpoAndRtoResp struct {

	// 任务ID
	JobId *string `json:"job_id,omitempty"`

	RpoInfo *RpoAndRtoInfo `json:"rpo_info,omitempty"`

	RtoInfo *RpoAndRtoInfo `json:"rto_info,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o QueryRpoAndRtoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryRpoAndRtoResp struct{}"
	}

	return strings.Join([]string{"QueryRpoAndRtoResp", string(data)}, " ")
}
