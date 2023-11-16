package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobConfigResponse Response Object
type ListJobConfigResponse struct {

	// 查询模板结果
	Result *[]CreateBuildJobRequestBody `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListJobConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobConfigResponse struct{}"
	}

	return strings.Join([]string{"ListJobConfigResponse", string(data)}, " ")
}
