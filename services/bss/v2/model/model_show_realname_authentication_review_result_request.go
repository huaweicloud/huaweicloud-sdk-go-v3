package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRealnameAuthenticationReviewResultRequest Request Object
type ShowRealnameAuthenticationReviewResultRequest struct {

	// 客户账号ID。您可以调用[查询客户列表](https://support.huaweicloud.com/api-bpconsole/mc_00021.html)接口获取customer_id。
	CustomerId string `json:"customer_id"`
}

func (o ShowRealnameAuthenticationReviewResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRealnameAuthenticationReviewResultRequest struct{}"
	}

	return strings.Join([]string{"ShowRealnameAuthenticationReviewResultRequest", string(data)}, " ")
}
