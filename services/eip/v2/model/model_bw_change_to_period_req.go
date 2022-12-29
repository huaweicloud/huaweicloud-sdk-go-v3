package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 云服务按需转包请求体
type BwChangeToPeriodReq struct {

	// 待按需转包带宽列表
	BandwidthIds []string `json:"bandwidth_ids"`

	ExtendParam *CreatePrePaidPublicipExtendParamOption `json:"extendParam"`
}

func (o BwChangeToPeriodReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BwChangeToPeriodReq struct{}"
	}

	return strings.Join([]string{"BwChangeToPeriodReq", string(data)}, " ")
}
