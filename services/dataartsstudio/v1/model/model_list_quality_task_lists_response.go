package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQualityTaskListsResponse Response Object
type ListQualityTaskListsResponse struct {

	// 错误码，如DQC.0000，请求处理成功
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListQualityTaskListsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQualityTaskListsResponse struct{}"
	}

	return strings.Join([]string{"ListQualityTaskListsResponse", string(data)}, " ")
}
