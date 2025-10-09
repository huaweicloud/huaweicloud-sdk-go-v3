package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccessPolicyRequest Request Object
type ListAccessPolicyRequest struct {

	// **参数解释：** 集群ID，仅返回与该集群相关的授权项。获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ClusterId *string `json:"cluster_id,omitempty"`
}

func (o ListAccessPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessPolicyRequest struct{}"
	}

	return strings.Join([]string{"ListAccessPolicyRequest", string(data)}, " ")
}
