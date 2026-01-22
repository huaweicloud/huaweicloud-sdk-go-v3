package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SessionPersistence struct {

	// **参数解释**：cookie名称。  **取值范围**：最大长度1024个字符。  [不支持该字段，请勿使用。](tag:hws_eu,hcso_dt)
	CookieName *string `json:"cookie_name,omitempty"`

	// **参数解释**：会话保持类型。  **取值范围**： - SOURCE_IP：基于源IP地址的会话保持，也就是根据客户端的源IP地址将流量持续转发到同一后端服务器（member）。 - HTTP_COOKIE：通过在HTTP响应中插入一个特殊的Cookie，实现会话保持。负载均衡会在后续请求中识别该Cookie，并将请求转发到对应的后端服务器。 - APP_COOKIE：基于应用层Cookie的会话保持。与HTTP_COOKIE不同的是，APP_COOKIE要求后端服务在响应中插入一个指定名称的Cookie，负载均衡器识别后，将后续请求转发到相同后端member。  [荷兰region不支持QUIC。](tag:dt) [不支持QUIC。](tag:tm)
	Type string `json:"type"`

	// **参数解释**：会话保持的时间。当type为APP_COOKIE时不生效。  **取值范围**： - 若pool的protocol为TCP、UDP和QUIC则范围为[1,60]（分钟），默认值1。 - 若pool的protocol为HTTP和HTTPS则范围为[1,1440]（分钟），默认值1440。  [荷兰region不支持QUIC。](tag:dt)
	PersistenceTimeout *int32 `json:"persistence_timeout,omitempty"`
}

func (o SessionPersistence) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SessionPersistence struct{}"
	}

	return strings.Join([]string{"SessionPersistence", string(data)}, " ")
}
