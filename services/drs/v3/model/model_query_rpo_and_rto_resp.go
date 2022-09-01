package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RPO和RTO信息体
type QueryRpoAndRtoResp struct {

	// 任务ID
	JobId *string `json:"job_id,omitempty" xml:"job_id"`

	RpoInfo *RpoAndRtoInfo `json:"rpo_info,omitempty" xml:"rpo_info"`

	RtoInfo *RpoAndRtoInfo `json:"rto_info,omitempty" xml:"rto_info"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg"`
}

func (o QueryRpoAndRtoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryRpoAndRtoResp struct{}"
	}

	return strings.Join([]string{"QueryRpoAndRtoResp", string(data)}, " ")
}
