package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBuildRecordBuildScriptResponse Response Object
type ShowBuildRecordBuildScriptResponse struct {

	// 构建脚本
	Result *string `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowBuildRecordBuildScriptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBuildRecordBuildScriptResponse struct{}"
	}

	return strings.Join([]string{"ShowBuildRecordBuildScriptResponse", string(data)}, " ")
}
