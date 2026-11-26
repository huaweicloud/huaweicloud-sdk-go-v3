package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RequestUrlRewriteEngine **参数解释：** 访问URL重写，当用户请求URL与CDN节点缓存资源的URL不一致时，可以通过访问URL重写功能重定向到目标URL **约束限制：** 不涉及
type RequestUrlRewriteEngine struct {

	// **参数解释：** 改写方式 **约束限制：** - regex为白名单controlled功能，非白名单用户不开放  **取值范围：** - simple: 精确改写 - regex: 正则改写  **默认取值：** simple: 精确改写
	RewriteType *string `json:"rewrite_type,omitempty"`

	// **参数解释：** 需要替换的URI **约束限制：** - 当rewrite_type为wildcard或regex时必填 - 当rewrite_type为regex时必须以“^/”开始，如：^/test  **取值范围：** - 1-512个字符 - 支持通配符\\*匹配，如：/test/\\*_/\\*.mp4 - 以正斜线（/）开头的URI，不含http(s)://头及域名  **默认取值：** 不涉及
	SourceUrl *string `json:"source_url,omitempty"`

	// **参数解释：** 重定向状态码 **约束限制：** 不涉及 **取值范围：** - 301: 资源被永久转移。 - 302: 由于不可预见的原因该页面暂不可用。 - 303: 用于PUT或POST请求完成之后进行页面跳转，防止由于页面刷新导致的操作的重复触发。 - 307: 由于不可预见的原因该页面暂不可用。  **默认取值：** 不涉及
	RedirectStatusCode *int32 `json:"redirect_status_code,omitempty"`

	// **参数解释：** 重定向URL **约束限制：** - 重定向后的URL，以正斜线（/）开头，不含http://头及域名，如：/test/index.html - 当匹配类型为全路径时，\"\\*\"可以用“$1”捕获，例如：匹配内容为/test/\\*.jpg，重定向URL配置为/newtest/$1.jpg，则用户请求/test/11.jpg时，$1捕获11，重定向后请求的URL为/newtest/11.jpg  **取值范围：** 不涉及 **默认取值：** 不涉及
	RedirectUrl string `json:"redirect_url"`

	// **参数解释：** 支持将客户端请求重定向到其他域名 **约束限制：** 不涉及 **取值范围：** - 1-255个字符 - 必须以http://或https://开头  **默认取值：** 不填时默认为当前域名
	RedirectHost *string `json:"redirect_host,omitempty"`

	// **参数解释：** 执行规则 **约束限制：** 不涉及 **取值范围：** - redirect: 如果请求的URL匹配了当前规则，该请求将被重定向到目标Path。执行完当前规则后，当存在其他配置规则时，会继续匹配剩余规则 - break: 如果请求的URL匹配了当前规则，请求将被改写为目标Path。执行完当前规则后，当存在其他配置规则时，将不再匹配剩余规则，此时不支持配置重定向Host和重定向状态码，返回状态码200  **默认取值：** 不涉及
	ExecutionMode string `json:"execution_mode"`

	// **参数解释：** 重定向URL是否携带请求参数 **约束限制：** 不涉及 **取值范围：** - on: 携带 - off: 不携带  **默认取值：** on: 携带
	InheritParam *string `json:"inherit_param,omitempty"`
}

func (o RequestUrlRewriteEngine) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RequestUrlRewriteEngine struct{}"
	}

	return strings.Join([]string{"RequestUrlRewriteEngine", string(data)}, " ")
}
