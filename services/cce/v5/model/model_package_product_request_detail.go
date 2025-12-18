package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PackageProductRequestDetail 订购套餐包的套餐包请求详情。
type PackageProductRequestDetail struct {

	// **参数解释：** 套餐包规格，通过套餐包列表接口获取。 **约束限制：** 不涉及 **取值范围：** 只支持通过套餐包列表接口获取的套餐包规格。 **默认取值：** 不涉及
	ResourceSpecCode string `json:"resource_spec_code"`

	// **参数解释：** 套餐包订购数量。 **约束限制：** 不涉及 **取值范围：** [1-10] **默认取值：** 0
	SubscriptionNum int32 `json:"subscription_num"`
}

func (o PackageProductRequestDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PackageProductRequestDetail struct{}"
	}

	return strings.Join([]string{"PackageProductRequestDetail", string(data)}, " ")
}
