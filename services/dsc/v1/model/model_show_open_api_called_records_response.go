package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowOpenApiCalledRecordsResponse struct {
	// 调用API总次数

	Total *int32 `json:"total,omitempty"`
	// 调用API成功次数

	Succeed *int32 `json:"succeed,omitempty"`
	// 调用API失败次数

	Failed *int32 `json:"failed,omitempty"`
	// API调用记录列表

	OpenapiCalledRecords *[]OpenApiCalledRecord `json:"openapi_called_records,omitempty"`
	// 获取下一页所需的标识符。

	NextMarker     *string `json:"next_marker,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowOpenApiCalledRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOpenApiCalledRecordsResponse struct{}"
	}

	return strings.Join([]string{"ShowOpenApiCalledRecordsResponse", string(data)}, " ")
}
