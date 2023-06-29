package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateReplSetNameRequest Request Object
type UpdateReplSetNameRequest struct {

	// 实例ID，可以调用“[查询实例列表和详情](x-wc://file=zh-cn_topic_0000001369935045.xml)”接口获取。如果未申请实例，可以调用“[创建实例](x-wc://file=zh-cn_topic_0000001369734929.xml)”接口创建。
	InstanceId string `json:"instance_id"`

	Body *MongoUpdateReplSetV3RequestBody `json:"body,omitempty"`
}

func (o UpdateReplSetNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateReplSetNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateReplSetNameRequest", string(data)}, " ")
}
