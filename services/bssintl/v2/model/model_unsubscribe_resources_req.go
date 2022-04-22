package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UnsubscribeResourcesReq struct {

	// 资源ID列表。 最大支持1次性输入10个资源ID，只能输入主资源ID。 哪些资源是主资源请根据“查询客户包年/包月资源列表”接口响应参数中的“is_main_resource”来标识。
	ResourceIds []string `json:"resource_ids"`

	// 退订类型，取值如下： 1：退订资源及其已续费周期。2：只退订资源已续费周期，不退订资源。
	UnsubscribeType int32 `json:"unsubscribe_type"`

	// 退订理由分类，取值如下： 1：产品不好用2：产品功能无法满足需求3：不会操作/操作过于复杂4：对服务不满意5：其他
	UnsubscribeReasonType *int32 `json:"unsubscribe_reason_type,omitempty"`

	// 退订原因，一般由客户输入。
	UnsubscribeReason *string `json:"unsubscribe_reason,omitempty"`
}

func (o UnsubscribeResourcesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnsubscribeResourcesReq struct{}"
	}

	return strings.Join([]string{"UnsubscribeResourcesReq", string(data)}, " ")
}
