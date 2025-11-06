package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteNotificationRequestBody struct {

	// 待删除的消息通知模板名称列表。
	Names []string `json:"names"`
}

func (o DeleteNotificationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNotificationRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteNotificationRequestBody", string(data)}, " ")
}
