package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResDatasourceResponse Response Object
type ShowResDatasourceResponse struct {
	Datasource *Datasources `json:"datasource,omitempty"`

	// 数据源相关任务详情。
	Jobs *[]Jobs `json:"jobs,omitempty"`

	// 是否成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 返回消息（请求成功时，不返回此字段）。
	Message *string `json:"message,omitempty"`

	// 错误码（请求成功时，不返回此字段）。
	ErrorCode      *string `json:"error_code,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowResDatasourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResDatasourceResponse struct{}"
	}

	return strings.Join([]string{"ShowResDatasourceResponse", string(data)}, " ")
}
