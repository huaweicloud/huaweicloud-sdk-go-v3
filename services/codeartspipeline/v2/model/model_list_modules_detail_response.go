package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListModulesDetailResponse Response Object
type ListModulesDetailResponse struct {
	Error *ErrorInfo `json:"error,omitempty"`

	// 插件的modules的具体信息集合
	Result map[string]ExtensionModuleList `json:"result,omitempty"`

	// 状态值
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListModulesDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListModulesDetailResponse struct{}"
	}

	return strings.Join([]string{"ListModulesDetailResponse", string(data)}, " ")
}
