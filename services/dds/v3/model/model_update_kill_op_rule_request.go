package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKillOpRuleRequest Request Object
type UpdateKillOpRuleRequest struct {

	// 实例ID，可以调用“[查询实例列表和详情](x-wc://file=zh-cn_topic_0000001369935045.xml)”接口获取。如果未申请实例，可以调用“[创建实例](x-wc://file=zh-cn_topic_0000001369734929.xml)”接口创建。
	InstanceId string `json:"instance_id"`

	Body *UpdateKillOpRuleRequestBody `json:"body,omitempty"`
}

func (o UpdateKillOpRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKillOpRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateKillOpRuleRequest", string(data)}, " ")
}
