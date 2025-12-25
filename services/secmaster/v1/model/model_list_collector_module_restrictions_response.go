package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCollectorModuleRestrictionsResponse Response Object
type ListCollectorModuleRestrictionsResponse struct {

	// 模板描述信息数组
	Body           *[]ShowTemplateFields `json:"body,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListCollectorModuleRestrictionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollectorModuleRestrictionsResponse struct{}"
	}

	return strings.Join([]string{"ListCollectorModuleRestrictionsResponse", string(data)}, " ")
}
