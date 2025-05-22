package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGeminiDbDualActiveRequest Request Object
type DeleteGeminiDbDualActiveRequest struct {

	// 参数解释 实例Id，可以调用[5.3.3查询实例列表和详情](x-wc://file=zh-cn_topic_0000001397299481.xml)接口获取。如果未申请实例，可以调用[5.3.1创建实例](x-wc://file=zh-cn_topic_0000001397139461.xml)接口创建。 约束限制 不涉及。 取值范围 不涉及。 默认取值 不涉及。
	InstanceId string `json:"instance_id"`
}

func (o DeleteGeminiDbDualActiveRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGeminiDbDualActiveRequest struct{}"
	}

	return strings.Join([]string{"DeleteGeminiDbDualActiveRequest", string(data)}, " ")
}
