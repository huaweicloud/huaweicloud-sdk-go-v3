package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowPlansResponse struct {
	// 项目下查询测试计划列表返回结构

	Body           *[]ShowPlansResponseBody `json:"body,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowPlansResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlansResponse struct{}"
	}

	return strings.Join([]string{"ShowPlansResponse", string(data)}, " ")
}
