package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmNotifyListRequest 查询告警消息列表请求参数模型。
type AlarmNotifyListRequest struct {

	// 页码。
	Page *int32 `json:"page,omitempty"`

	// 每页数量。
	PageSize *int32 `json:"page_size,omitempty"`

	// 告警事件id。
	AlarmDataId int32 `json:"alarm_data_id"`

	// region英文名称。
	Region string `json:"region"`
}

func (o AlarmNotifyListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmNotifyListRequest struct{}"
	}

	return strings.Join([]string{"AlarmNotifyListRequest", string(data)}, " ")
}
