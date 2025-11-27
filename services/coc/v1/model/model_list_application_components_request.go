package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplicationComponentsRequest Request Object
type ListApplicationComponentsRequest struct {

	// 应用id。
	ApplicationId *string `json:"application_id,omitempty"`

	// 分页参数，上一页请求最后一个id。
	Marker *string `json:"marker,omitempty"`

	// 最大的返回数量。
	Limit int32 `json:"limit"`
}

func (o ListApplicationComponentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationComponentsRequest struct{}"
	}

	return strings.Join([]string{"ListApplicationComponentsRequest", string(data)}, " ")
}
