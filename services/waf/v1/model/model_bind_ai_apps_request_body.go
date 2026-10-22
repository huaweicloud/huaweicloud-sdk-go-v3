package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindAiAppsRequestBody 绑定 AI 应用请求
type BindAiAppsRequestBody struct {

	// 绑定的 AI 应用 ID 列表
	AppIds *[]string `json:"app_ids,omitempty"`
}

func (o BindAiAppsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindAiAppsRequestBody struct{}"
	}

	return strings.Join([]string{"BindAiAppsRequestBody", string(data)}, " ")
}
