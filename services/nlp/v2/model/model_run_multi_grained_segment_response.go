package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunMultiGrainedSegmentResponse struct {
	// 多粒度分词结果列表。调用失败时无此字段。

	Result *[]PostMultiGainedSegmentResponseItem `json:"result,omitempty"`
	// 调用失败时的错误码，具体请参见错误码。调用成功时无此字段。

	ErrorCode *string `json:"error_code,omitempty"`
	// 调用失败时的错误信息。调用成功时无此字段。

	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunMultiGrainedSegmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunMultiGrainedSegmentResponse struct{}"
	}

	return strings.Join([]string{"RunMultiGrainedSegmentResponse", string(data)}, " ")
}
