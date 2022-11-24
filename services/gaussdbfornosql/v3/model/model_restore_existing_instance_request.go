package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RestoreExistingInstanceRequest struct {

	// 实例Id，可以调用[5.3.3 查询实例列表和详情](x-wc://file=zh-cn_topic_0000001397299481.xml)接口获取。如果未申请实例，可以调用[5.3.1 创建实例](x-wc://file=zh-cn_topic_0000001397139461.xml)接口创建。
	InstanceId string `json:"instance_id"`

	Body *RestoreRequestBody `json:"body,omitempty"`
}

func (o RestoreExistingInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreExistingInstanceRequest struct{}"
	}

	return strings.Join([]string{"RestoreExistingInstanceRequest", string(data)}, " ")
}
