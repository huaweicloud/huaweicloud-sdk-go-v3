package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyInferIntranetConnectionsRequest Request Object
type ModifyInferIntranetConnectionsRequest struct {

	// 参数解释： 内网接入id。id可以根据[查询当前租户的内网接入申请列表](ListInferIntranetConnectionApplications.xml)返回body的id字段得到。 约束限制： 不涉及。 取值范围： 不涉及。 默认取值： 不涉及。
	Id string `json:"id"`

	Body *IntranetConnectionModifyRequest `json:"body,omitempty"`
}

func (o ModifyInferIntranetConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyInferIntranetConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ModifyInferIntranetConnectionsRequest", string(data)}, " ")
}
