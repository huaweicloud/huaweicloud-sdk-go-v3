package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListResDatasourcesResponse struct {

	// 数据源详情列表。
	Datasources *[]Datasources `json:"datasources,omitempty"`

	// 是否成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 返回消息（请求成功时，不返回此字段）。
	Message *string `json:"message,omitempty"`

	// 错误码（请求成功时，不返回此字段）。
	ErrorCode      *string `json:"error_code,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListResDatasourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResDatasourcesResponse struct{}"
	}

	return strings.Join([]string{"ListResDatasourcesResponse", string(data)}, " ")
}
