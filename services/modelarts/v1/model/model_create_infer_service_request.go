package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInferServiceRequest Request Object
type CreateInferServiceRequest struct {

	// **参数解释：** 服务提供者的domain级或project级Token，创建服务携带该请求头时，系统将解析该token并将账号id保存为服务的提供者即provider，该服务将被系统保护，仅携带该提供者的domain级或project级Token的更新操作允许执行。[通过调用IAM服务获取用户Token接口获取响应消息头中X-Subject-Token的值。](tag:hws,hws_hk)获取方法请参见[[获取IAM用户Token（使用密码）](modelarts_03_0004.xml)](tag:hws,hws_hk)[[获取Token](modelarts_03_0015.xml)](tag:hcs,hcs_sm)。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	XAuthTokenProvider *string `json:"X-Auth-Token-Provider,omitempty"`

	Body *ServiceCreateRequest `json:"body,omitempty"`
}

func (o CreateInferServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInferServiceRequest struct{}"
	}

	return strings.Join([]string{"CreateInferServiceRequest", string(data)}, " ")
}
