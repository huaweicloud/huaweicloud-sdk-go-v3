package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostGroupBaseInfosResponse Response Object
type ListHostGroupBaseInfosResponse struct {

	// 请求成功失败状态
	Status *string `json:"status,omitempty"`

	// 总数量
	Total *int32 `json:"total,omitempty"`

	// 环境基本信息列表
	Result         *[]EnvironmentBaseInfo `json:"result,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListHostGroupBaseInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostGroupBaseInfosResponse struct{}"
	}

	return strings.Join([]string{"ListHostGroupBaseInfosResponse", string(data)}, " ")
}
