package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExchangeInstanceVipProxyIpResponse Response Object
type ExchangeInstanceVipProxyIpResponse struct {

	// **参数解释**：  实例与数据库代理IP互换的任务ID。  **取值范围**：  不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExchangeInstanceVipProxyIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExchangeInstanceVipProxyIpResponse struct{}"
	}

	return strings.Join([]string{"ExchangeInstanceVipProxyIpResponse", string(data)}, " ")
}
