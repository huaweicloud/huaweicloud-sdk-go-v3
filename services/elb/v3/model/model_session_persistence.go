package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 会话持久性对象。
type SessionPersistence struct {
	// cookie名称。  格式：仅支持字母、数字、中划线(-)、下划线(_)和点号(.)。  [使用说明： - 只有当type为APP_COOKIE时才有效。](tag:hws,hws_hk,ocb,tlf,ctc,hcso,sbc,g42,tm,cmcc,hk-g42,dt,dt_test)  [不支持该字段，请勿使用。](tag:hcso_dt)

	CookieName *string `json:"cookie_name,omitempty"`
	// 会话保持类型。 取值范围：SOURCE_IP、HTTP_COOKIE、APP_COOKIE。 使用说明：  - 当pool的protocol为TCP、UDP，无论type取值如何，都会被忽略，会话保持只按SOURCE_IP生效；  [- 当pool的protocol为HTTP、HTTPS时，如果是独享型负载均衡器的pool，则type只能为HTTP_COOKIE，其他取值会话保持失效。如果是共享型负载均衡器的pool，则type可以为HTTP_COOKIE和APP_COOKIE，其他取值会话保持失效。 ](tag:hws,hws_hk,ocb,tlf,ctc,hcso,sbc,g42,tm,cmcc,hk-g42,dt,dt_test) [- 当pool的protocol为HTTP、HTTPS时，type只能为HTTP_COOKIE，其他取值会话保持失效。](tag:hcso_dt)  [不支持APP_COOKIE，请勿使用。](tag:hcso_dt)

	Type string `json:"type"`
	// 会话保持的时间。[当type为APP_COOKIE时不生效。](tag:hws,hws_hk,ocb,tlf,ctc,hcso,sbc,g42,tm,cmcc,hk-g42,dt,dt_test)  适用范围：如果pool的protocol为TCP、UDP和QUIC则范围为[1,60]（分钟），默认值1；如果pool的protocol为HTTP和HTTPS则范围为[1,1440]（分钟），默认值1440。

	PersistenceTimeout *int32 `json:"persistence_timeout,omitempty"`
}

func (o SessionPersistence) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SessionPersistence struct{}"
	}

	return strings.Join([]string{"SessionPersistence", string(data)}, " ")
}
