package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddReadonlyNodeRequestBody struct {

	// 资源规格编码。获取方法请参见[查询数据库规格](x-wc://file=zh-cn_topic_0000001321087266.xml)中参数“spec_code”的值。  示例：dds.mongodb.c6.xlarge.2.shard
	SpecCode string `json:"spec_code"`

	// 待新增只读节点个数。 取值范围：1-5。
	Num string `json:"num"`

	// 同步延迟时间。取值范围：0~1200毫秒。默认取值为0。
	Delay *int32 `json:"delay,omitempty"`

	// 扩容包年包月实例的存储容量时可指定，表示是否自动从账户中支付，此字段不影响自动续订的支付方式。 - true，表示自动从账户中支付。 - false，表示手动从账户中支付，默认为该方式。
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
}

func (o AddReadonlyNodeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddReadonlyNodeRequestBody struct{}"
	}

	return strings.Join([]string{"AddReadonlyNodeRequestBody", string(data)}, " ")
}
