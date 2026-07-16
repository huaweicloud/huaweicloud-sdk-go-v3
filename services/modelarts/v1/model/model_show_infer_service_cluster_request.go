package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInferServiceClusterRequest Request Object
type ShowInferServiceClusterRequest struct {

	// **参数解释：** 资源池ID，查询指定资源池下的服务，默认不过滤。可通过[查询资源池列表](ShowPool.xml)获取。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Id string `json:"id"`

	// **参数解释：** 用户Token。[通过调用IAM服务获取用户Token接口获取响应消息头中X-Subject-Token的值。](tag:hws,hws_hk)获取方法请参见[[获取IAM用户Token（使用密码）](modelarts_03_0004.xml)](tag:hws,hws_hk)[[获取Token](modelarts_03_0015.xml)](tag:hcs,hcs_sm)。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	XUserToken *string `json:"X-User-Token,omitempty"`
}

func (o ShowInferServiceClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInferServiceClusterRequest struct{}"
	}

	return strings.Join([]string{"ShowInferServiceClusterRequest", string(data)}, " ")
}
