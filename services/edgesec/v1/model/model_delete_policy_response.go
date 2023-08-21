package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePolicyResponse Response Object
type DeletePolicyResponse struct {

	// 防护策略id
	Id *string `json:"id,omitempty"`

	// 防护策略名
	Name *string `json:"name,omitempty"`

	// Web基础防护等级   - 1 : 宽松，防护粒度较粗，只拦截攻击特征比较明显的请求。当误报情况较多的场景下，建议选择“宽松”模式。   - 2：中等，默认为“中等”防护模式，满足大多数场景下的Web防护需求。   - 3：严格，防护粒度最精细，可以拦截具有复杂的绕过特征的攻击请求，例如jolokia网络攻击、探测CGI漏洞、探测 Druid SQL注入攻击
	Level *int32 `json:"level,omitempty"`

	// 精准防护中的检测模式。   - false：短路检测，当用户的请求符合精准防护中的拦截条件时，便立刻终止检测，进行拦截   - true ：全检测，请求符合精准防护中的拦截条件时，全检测不会立即拦截，会继续执行其他防护的检测，最后进行拦截。
	FullDetection *bool `json:"full_detection,omitempty"`

	RobotAction *Action `json:"robot_action,omitempty"`

	Action *WafPolicyAction `json:"action,omitempty"`

	Options *PolicyOption `json:"options,omitempty"`

	// 智能访问控制防护项相关配置信息，目前该特性还处于公测阶段，只有部分局点支持该特性
	ModulexOptions map[string]interface{} `json:"modulex_options,omitempty"`

	// 与防护策略绑定的防护的域名id数组
	Hosts *[]string `json:"hosts,omitempty"`

	// 与防护策略绑定的防护的域名信息数组，相对于hosts字段，包含更详细的域名信息
	BindHost *[]BindHost `json:"bind_host,omitempty"`

	// 扩展字段，用于存放Web基础防护中一些开关配置等信息
	Extend map[string]string `json:"extend,omitempty"`

	// 创建防护策略的时间
	Timestamp      *int64 `json:"timestamp,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o DeletePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePolicyResponse struct{}"
	}

	return strings.Join([]string{"DeletePolicyResponse", string(data)}, " ")
}
