package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachBatchPublicIpResponse Response Object
type DetachBatchPublicIpResponse struct {

	// 弹性公网IP对象
	Publicips *[]BatchPublicipResp `json:"publicips,omitempty"`

	// 本次请求编号
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DetachBatchPublicIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachBatchPublicIpResponse struct{}"
	}

	return strings.Join([]string{"DetachBatchPublicIpResponse", string(data)}, " ")
}
