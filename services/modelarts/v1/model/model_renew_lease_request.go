package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RenewLeaseRequest Request Object
type RenewLeaseRequest struct {

	// **参数解释**：续订时长，推荐该参数在leaseReq中配置，若请求参数中包含duration，则忽略leaseReq的值，且实例自动停止类别为定时停止。(单位:毫秒)。 **约束限制**：不涉及。 **取值范围**：3600000-259200000。 **默认取值**：3600000。
	Duration *int64 `json:"duration,omitempty"`

	// **参数解释**：Notebook实例ID。ID格式为通用唯一识别码（Universally Unique Identifier，简称UUID），可通过调用[[查询Notebook实例列表接口](https://support.huaweicloud.com/api-modelarts/ListNotebooks.html#section0)](tag:hc)[[查询Notebook实例列表接口](https://support.huaweicloud.com/intl/zh-cn/api-modelarts/ListNotebooks.html#section0)](tag:hk)获取。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Id string `json:"id"`

	Body *LeaseReq `json:"body,omitempty"`
}

func (o RenewLeaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RenewLeaseRequest struct{}"
	}

	return strings.Join([]string{"RenewLeaseRequest", string(data)}, " ")
}
