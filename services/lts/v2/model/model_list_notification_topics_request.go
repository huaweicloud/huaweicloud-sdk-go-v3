package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNotificationTopicsRequest Request Object
type ListNotificationTopicsRequest struct {

	// 查询游标，初始传入0，后续从上一次的返回值中获取
	Offset int32 `json:"offset"`

	// 每页数据量，最大值为100
	Limit int32 `json:"limit"`
}

func (o ListNotificationTopicsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationTopicsRequest struct{}"
	}

	return strings.Join([]string{"ListNotificationTopicsRequest", string(data)}, " ")
}
