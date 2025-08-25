package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCcspInstanceInfoResponse Response Object
type ShowCcspInstanceInfoResponse struct {

	// 满足条件的实例总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 实例列表
	Result         *[]CcspInstanceInfo `json:"result,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowCcspInstanceInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCcspInstanceInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowCcspInstanceInfoResponse", string(data)}, " ")
}
