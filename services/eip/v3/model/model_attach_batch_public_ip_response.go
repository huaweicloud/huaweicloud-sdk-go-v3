package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachBatchPublicIpResponse Response Object
type AttachBatchPublicIpResponse struct {

	// 弹性公网IP对象
	Publicips *[]BatchPublicipResp `json:"publicips,omitempty"`

	// 本次请求编号
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AttachBatchPublicIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachBatchPublicIpResponse struct{}"
	}

	return strings.Join([]string{"AttachBatchPublicIpResponse", string(data)}, " ")
}
