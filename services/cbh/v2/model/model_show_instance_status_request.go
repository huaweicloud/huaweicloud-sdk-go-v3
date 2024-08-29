package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceStatusRequest Request Object
type ShowInstanceStatusRequest struct {

	// 云堡垒机实例ID，使用UUID格式表示。  获取方法详见用户指南里面的实例\"查看实例详情\"
	ServerId string `json:"server_id"`
}

func (o ShowInstanceStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceStatusRequest", string(data)}, " ")
}
