package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlertTemplatesRequest Request Object
type ListAlertTemplatesRequest struct {

	// 服务id
	ServiceId string `json:"service_id"`

	// 模板名称
	Name *string `json:"name,omitempty"`

	// 当前页数
	PageNum *int32 `json:"pageNum,omitempty"`

	// 每页数量
	PageSize *int32 `json:"pageSize,omitempty"`
}

func (o ListAlertTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlertTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListAlertTemplatesRequest", string(data)}, " ")
}
