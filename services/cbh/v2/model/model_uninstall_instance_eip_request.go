package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UninstallInstanceEipRequest Request Object
type UninstallInstanceEipRequest struct {

	// 云堡垒机实例ID，使用UUID格式表示。  获取方法详见用户指南里面的实例\"查看实例详情\"
	ServerId string `json:"server_id"`

	Body *OperateEipRequestBody `json:"body,omitempty"`
}

func (o UninstallInstanceEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UninstallInstanceEipRequest struct{}"
	}

	return strings.Join([]string{"UninstallInstanceEipRequest", string(data)}, " ")
}
