package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAssociateEnvironmentsInfosResponse Response Object
type ListAssociateEnvironmentsInfosResponse struct {

	// 请求成功失败状态
	Status *string `json:"status,omitempty"`

	// 关联环境总数量
	Total *int32 `json:"total,omitempty"`

	// 环境信息列表
	Result         *[]EnvironmentInfo `json:"result,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListAssociateEnvironmentsInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssociateEnvironmentsInfosResponse struct{}"
	}

	return strings.Join([]string{"ListAssociateEnvironmentsInfosResponse", string(data)}, " ")
}
