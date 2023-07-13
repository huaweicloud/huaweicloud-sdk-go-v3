package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNotificationTemplatesRequest Request Object
type ListNotificationTemplatesRequest struct {

	// 账号id，获取方式请参见：获取账号ID、项目ID、日志组ID、日志流ID（https://support.huaweicloud.com/api-lts/lts_api_0006.html）。
	DomainId string `json:"domain_id"`

	// 查询游标，初始传入0，后续从上一次的返回值中获取
	Offset *int32 `json:"offset,omitempty"`

	// 每页数据量，最大值为100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListNotificationTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListNotificationTemplatesRequest", string(data)}, " ")
}
