package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSessionRequestBodyV2 **参数解释**：创建session请求体。 **约束限制**：不涉及。
type CreateSessionRequestBodyV2 struct {

	// **参数解释**：Endpoint名称。 **约束限制**：不涉及。 **取值范围**：长度为1~128个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	EndpointName string `json:"endpoint_name"`

	// **参数解释**：委托名称。 **约束限制**：具体请参见[[配置Aura委托](https://support.huaweicloud.com/usermanual-aidatalake/aidatalake_01_0078.html#section4)](tag:HEC)[[配置Aura委托](https://support.huaweicloud.com/intl/zh-cn/usermanual-aidatalake/aidatalake_01_0078.html#section4)](tag:hws_hk)。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AgencyName *string `json:"agency_name,omitempty"`

	CatalogContext *CatalogContext `json:"catalog_context"`

	// **参数解释**：创建session排队等待时间，单位：秒。 **约束限制**：不涉及。 **取值范围**：300~86400。 **默认取值**：不涉及。
	WaitTimeout *int32 `json:"wait_timeout,omitempty"`

	// **参数解释**：资源不足时重试次数。 **约束限制**：不涉及。 **取值范围**：1~2147483647。 **默认取值**：不涉及。
	MaxRetryTimes *int32 `json:"max_retry_times,omitempty"`

	// **参数解释**：优先级（1~9，9为最高优先级）。 **约束限制**：不涉及。 **取值范围**：1~9。 **默认取值**：5。
	Priority *int32 `json:"priority,omitempty"`

	// **参数解释**：会话维度指定的 CN（计算节点）CPU 资源规格，单位：毫核（millicores）。 **约束限制**：可选；不传时回退到全局配置 getCnCpuLimit； INNER_SESSION/维护语句场景不消费该字段，资源规格继承父会话 CN； 上限 48000 毫核（48 核），超上限返回 400 INVALID_PARAMETER。 **取值范围**：(0, 48000] 毫核。 **默认取值**：不传时取全局配置兜底值（如 4000 毫核）。
	Cpu *int64 `json:"cpu,omitempty"`

	// **参数解释**：会话维度指定的 CN（计算节点）Memory 资源规格，单位：MB。 **约束限制**：可选；不传时回退到全局配置 getCnMemoryLimit； INNER_SESSION/维护语句场景不消费该字段，资源规格继承父会话 CN； 上限 393216 MB（384 GB），超上限返回 400 INVALID_PARAMETER。 **取值范围**：(0, 393216] MB。 **默认取值**：不传时取全局配置兜底值（如 16384 MB）。
	Memory *int64 `json:"memory,omitempty"`
}

func (o CreateSessionRequestBodyV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSessionRequestBodyV2 struct{}"
	}

	return strings.Join([]string{"CreateSessionRequestBodyV2", string(data)}, " ")
}
