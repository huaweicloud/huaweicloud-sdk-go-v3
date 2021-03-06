package model

import (
	"encoding/json"

	"strings"
)

// 会话持久性对象。
type SessionPersistence struct {
	// cookie名称。 只有当type为APP_COOKIE时才支持。 格式要求：仅支持字母数字-_.

	CookieName *string `json:"cookie_name,omitempty"`
	// 描述：类型，可以为SOURCE_IP、HTTP_COOKIE、APP_COOKIE。   约束：   1、当pool的protocol为TCP、UDP、QUIC时，只按SOURCE_IP生效；   2、当pool的protocol为HTTP、HTTPS时，只按HTTP_COOKIE、APP_COOKIE生效。

	Type string `json:"type"`
	// 会话保持的时间。当type为APP_COOKIE时不生效。 适用范围：如果pool的protocol为TCP、UDP和QUIC则范围为[1,60]（分钟），默认值1；如果pool的protocol为HTTP和HTTPS则范围为[1,1440]（分钟），默认值1440。

	PersistenceTimeout *int32 `json:"persistence_timeout,omitempty"`
}

func (o SessionPersistence) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SessionPersistence struct{}"
	}

	return strings.Join([]string{"SessionPersistence", string(data)}, " ")
}
