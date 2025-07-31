package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AutoRenewalResourcesReq struct {

	// |参数名称：自动续费次数| |参数的约束及描述：该参数非必填，范围限制：0-99，0代表不限制次数。 首次开通自动续费，此参数不携带或携带值为null时，默认为不限制次数 已开通自动续费，重置自动续费次数时此参数必填，否则不做处理，不进行修改|
	AutoRenewTimes *int32 `json:"auto_renew_times,omitempty"`
}

func (o AutoRenewalResourcesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoRenewalResourcesReq struct{}"
	}

	return strings.Join([]string{"AutoRenewalResourcesReq", string(data)}, " ")
}
