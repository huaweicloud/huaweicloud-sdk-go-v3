package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAppQuotaBoundAppsResponse struct {
	// 本次返回的列表长度

	Size int32 `json:"size"`
	// 满足条件的记录数

	Total int64 `json:"total"`
	// 客户端应用列表

	Apps           *[]AppQuotaAppInfo `json:"apps,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListAppQuotaBoundAppsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppQuotaBoundAppsResponse struct{}"
	}

	return strings.Join([]string{"ListAppQuotaBoundAppsResponse", string(data)}, " ")
}
