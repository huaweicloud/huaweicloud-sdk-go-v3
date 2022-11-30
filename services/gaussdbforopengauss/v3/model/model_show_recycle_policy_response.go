package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowRecyclePolicyResponse struct {

	// 已删除实例保留天数，可设置范围为1~7天。 - 取值1~7，设置已删除实例的保留天数为该值。
	RetentionPeriodInDays *string `json:"retention_period_in_days,omitempty"`
	HttpStatusCode        int     `json:"-"`
}

func (o ShowRecyclePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecyclePolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowRecyclePolicyResponse", string(data)}, " ")
}
