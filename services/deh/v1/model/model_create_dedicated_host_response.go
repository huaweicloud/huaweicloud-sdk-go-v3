package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDedicatedHostResponse Response Object
type CreateDedicatedHostResponse struct {

	// 分配专属主机的Job ID。
	JobId *string `json:"job_id,omitempty"`

	// 分配包周期专属主机的订单ID。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDedicatedHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDedicatedHostResponse struct{}"
	}

	return strings.Join([]string{"CreateDedicatedHostResponse", string(data)}, " ")
}
