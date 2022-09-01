package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchCheckJobsResponse struct {

	// 预检查响应体
	Results *[]PostPreCheckResp `json:"results,omitempty" xml:"results"`

	// 总数
	Count          *int32 `json:"count,omitempty" xml:"count"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchCheckJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCheckJobsResponse struct{}"
	}

	return strings.Join([]string{"BatchCheckJobsResponse", string(data)}, " ")
}
