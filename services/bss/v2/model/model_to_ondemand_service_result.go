package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ToOndemandServiceResult struct {

	// |参数名称：转按需的服务实例id/根资源id| |参数约束以及描述：转按需的服务实例id/根资源id。对应请求体中的资源id。|
	ResourceId *string `json:"resource_id,omitempty"`

	// |参数名称：转按需结果| |参数约束以及描述：转按需结果。SUCCESS：成功；AUDIT：审核中；FAIL：转按需失败。|
	Result *string `json:"result,omitempty"`

	// |参数名称：状态码| |参数约束以及描述：状态码，result=FAIL时，必填|
	ErrorCode *string `json:"error_code,omitempty"`

	// |参数名称：错误描述信息| |参数约束以及描述：错误描述信息，result=FAIL时，必填|
	ErrorMsg *string `json:"error_msg,omitempty"`

	// |参数名称：订单Id| |参数约束以及描述：订单Id。result= SUCCESS、 AUDIT时，必填；其他为空。|
	OrderId *string `json:"order_id,omitempty"`
}

func (o ToOndemandServiceResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ToOndemandServiceResult struct{}"
	}

	return strings.Join([]string{"ToOndemandServiceResult", string(data)}, " ")
}
