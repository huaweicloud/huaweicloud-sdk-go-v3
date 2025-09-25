package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateEsElbRequestBody struct {

	// 打开或关闭es负载均衡器。 - true：开启。 - false：关闭。
	Enable bool `json:"enable"`

	// 委托名称。
	Agency *string `json:"agency,omitempty"`

	// 负载均衡器id。
	ElbId *string `json:"elb_id,omitempty"`

	// 负载均衡器类型。当用于可观测集群时：不需要配置此参数。
	Type *string `json:"type,omitempty"`
}

func (o UpdateEsElbRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEsElbRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEsElbRequestBody", string(data)}, " ")
}
