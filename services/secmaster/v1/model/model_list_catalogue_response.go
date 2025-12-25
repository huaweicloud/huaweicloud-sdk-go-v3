package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCatalogueResponse Response Object
type ListCatalogueResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 数据列表
	Data *[]CatalogueDetailInfo `json:"data,omitempty"`

	// 响应信息
	Message *string `json:"message,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	// 是否响应成功
	Success *bool `json:"success,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListCatalogueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCatalogueResponse struct{}"
	}

	return strings.Join([]string{"ListCatalogueResponse", string(data)}, " ")
}
