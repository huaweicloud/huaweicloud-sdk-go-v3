package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 后端云服务器组的会话持久性。 当开启会话保持后，在一定时间内，来自同一客户端的请求会发送到同一个后端云服务器上。 当会话保持关闭时，该字段取值为null。
type SessionPersistence struct {
	// 会话保持的类型。SOURCE_IP：根据请求的源IP，将同一IP的请求发送到同一个后端云服务器上。HTTP_COOKIE：客户端第一次发送请求时，负载均衡器自动生成cookie并将该cookie插入响应消息中，后续请求会发送到处理第一个请求的后端云服务器上。APP_COOKIE：客户端第一次发送请求时，后端服务器生成cookie并将该cookie插入响应消息中，后续请求会发送到处理第一个请求的后端云服务器上。当后端云服务器的protocol为TCP时，只按SOURCE_IP生效当后端云服务器的protocol为HTTP时，只按HTTP_COOKIE或APP_COOKIE生效

	Type SessionPersistenceType `json:"type"`
	// cookie的名称。只有当会话保持的类型是APP_COOKIE时可以指定。

	CookieName *string `json:"cookie_name,omitempty"`
	// 会话保持的超时时间。取值范围：[1,60]（分钟）：当后端云服务器的protocol为TCP、UDP时[1,1440]（分钟）：当后端云服务器的protocol为HTTP时。当type为APP_COOKIE时该字段不生效。

	PersistenceTimeout *int32 `json:"persistence_timeout,omitempty"`
}

func (o SessionPersistence) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SessionPersistence struct{}"
	}

	return strings.Join([]string{"SessionPersistence", string(data)}, " ")
}

type SessionPersistenceType struct {
	value string
}

type SessionPersistenceTypeEnum struct {
	SOURCE_IP   SessionPersistenceType
	HTTP_COOKIE SessionPersistenceType
	APP_COOKIE  SessionPersistenceType
}

func GetSessionPersistenceTypeEnum() SessionPersistenceTypeEnum {
	return SessionPersistenceTypeEnum{
		SOURCE_IP: SessionPersistenceType{
			value: "SOURCE_IP",
		},
		HTTP_COOKIE: SessionPersistenceType{
			value: "HTTP_COOKIE",
		},
		APP_COOKIE: SessionPersistenceType{
			value: "APP_COOKIE",
		},
	}
}

func (c SessionPersistenceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SessionPersistenceType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
