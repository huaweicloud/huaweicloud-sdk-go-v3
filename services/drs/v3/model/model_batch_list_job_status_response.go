package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchListJobStatusResponse struct {

	// 任务状态信息
	Results *[]QueryJobStatusResp `json:"results,omitempty" xml:"results"`

	// 返回任务数量
	Count          *int32 `json:"count,omitempty" xml:"count"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchListJobStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListJobStatusResponse struct{}"
	}

	return strings.Join([]string{"BatchListJobStatusResponse", string(data)}, " ")
}
