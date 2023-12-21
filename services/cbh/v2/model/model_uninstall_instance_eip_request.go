package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UninstallInstanceEipRequest Request Object
type UninstallInstanceEipRequest struct {

	// 云堡垒机实例ID，使用UUID格式表示。  [实例ID获取方式](https://support.huaweicloud.com/usermanual-cbh/cbh_02_1003.html)。
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
