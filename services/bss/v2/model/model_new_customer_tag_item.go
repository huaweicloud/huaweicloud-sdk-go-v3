package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NewCustomerTagItem struct {

	// 客户ID。
	CustomerId *string `json:"customer_id,omitempty"`

	// 新客标签类型。 SMB：SMB新客标签。
	CustomerTagType *string `json:"customer_tag_type,omitempty"`

	// 新客标签。 Y：合格新客N：非新客UNDETERMINED：未达标新客，即有新客资格但尚未成为新客
	NewCustomerTag *string `json:"new_customer_tag,omitempty"`

	// 生效月份。 格式为YYYY-MM。
	EffectiveTime *string `json:"effective_time,omitempty"`

	// 失效月份。 格式为YYYY-MM。
	ExpireTime *string `json:"expire_time,omitempty"`
}

func (o NewCustomerTagItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NewCustomerTagItem struct{}"
	}

	return strings.Join([]string{"NewCustomerTagItem", string(data)}, " ")
}
