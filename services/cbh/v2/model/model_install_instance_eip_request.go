package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstallInstanceEipRequest Request Object
type InstallInstanceEipRequest struct {

	// 云堡垒机实例ID，使用UUID格式表示。  [实例ID获取方式](https://support.huaweicloud.com/usermanual-cbh/cbh_02_1003.html)。
	ServerId string `json:"server_id"`

	Body *OperateEipRequestBody `json:"body,omitempty"`
}

func (o InstallInstanceEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstallInstanceEipRequest struct{}"
	}

	return strings.Join([]string{"InstallInstanceEipRequest", string(data)}, " ")
}
