package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowProjectTraceResponse struct {

	// 数据对象（目录，文件）总数量
	Count *int64 `json:"count,omitempty"`

	// 数据对象列表
	Datas *[]TraceDataRsp `json:"datas,omitempty"`

	// 桶存量
	BucketSize     *int64 `json:"bucket_size,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowProjectTraceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectTraceResponse struct{}"
	}

	return strings.Join([]string{"ShowProjectTraceResponse", string(data)}, " ")
}
