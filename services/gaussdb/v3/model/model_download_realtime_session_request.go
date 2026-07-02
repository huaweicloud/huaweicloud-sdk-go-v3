package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadRealtimeSessionRequest Request Object
type DownloadRealtimeSessionRequest struct {

	// 参数解释：  实例ID，此参数是实例的唯一标识。  获取方法请参见[查询实例列表](https://support.huaweicloud.com/api-taurusdb/ShowMysqlInstanceList.html)。  约束限制：  不涉及。  取值范围：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  默认取值：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  节点ID。 获取方法请参见[查询实例详情信息](https://support.huaweicloud.com/api-taurusdb/ShowGaussMySqlInstanceInfoUnifyStatus.html)。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，前面为UUID，后缀为no07，长度为36个字符。  **默认取值**：  不涉及。
	NodeId string `json:"node_id"`

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**：  收集全部实时会话信息任务ID。  获取方法请参见[收集全部实时会话信息](https://support.huaweicloud.com/api-taurusdb/CollectRealtimeSession.html)。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	RequestId string `json:"request_id"`
}

func (o DownloadRealtimeSessionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadRealtimeSessionRequest struct{}"
	}

	return strings.Join([]string{"DownloadRealtimeSessionRequest", string(data)}, " ")
}
