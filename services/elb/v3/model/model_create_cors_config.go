package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCorsConfig **参数解释**：转发策略跨域的配置。  **约束限制**：不涉及
type CreateCorsConfig struct {

	// **参数解释**：允许的访问来源列表。  **约束限制**： - 若配置了通配符（*），则不能配其他值。 - 不支持IP地址。 - 不能配置重复的值。  **取值范围**：单个列表项的取值 - 通配符（*），表示匹配所有来源。 - 除通配符外，其他值只能是以http://或者https://开头（小写），后面加域名。可以是具体域名或一级泛域名，可以包含端口或不包含端口。例：http://_*.test.abc.example.com:80  **默认取值**：不涉及
	AllowOrigin *[]string `json:"allow_origin,omitempty"`

	// **参数解释**：选择跨域访问时允许的 HTTP 方法。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	AllowMethods *[]string `json:"allow_methods,omitempty"`

	// **参数解释**：允许跨域的 Header 列表。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	AllowHeaders *[]string `json:"allow_headers,omitempty"`

	// **参数解释**：允许暴露的Header列表。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	ExposeHeaders *[]string `json:"expose_headers,omitempty"`

	// **参数解释**：是否允许携带凭证信息。  **约束限制**：不涉及  **取值范围**：true 是，false 否。  **默认取值**：不涉及
	AllowCredentials *bool `json:"allow_credentials,omitempty"`

	// **参数解释**：预检请求在浏览器的最大缓存时间，单位：秒。  **约束限制**：不涉及  **取值范围**：-1~172800  **默认取值**：不涉及
	MaxAge *int64 `json:"max_age,omitempty"`
}

func (o CreateCorsConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCorsConfig struct{}"
	}

	return strings.Join([]string{"CreateCorsConfig", string(data)}, " ")
}
