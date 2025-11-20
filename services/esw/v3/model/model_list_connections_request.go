package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConnectionsRequest Request Object
type ListConnectionsRequest struct {

	// - 参数解释：ESW实例的唯一标识。 - 约束限制：带“-”的UUID格式。 - 取值范围：不涉及。 - 默认取值：不涉及。
	InstanceId string `json:"instance_id"`

	// - 参数解释：二层连接资源ID。二层连接创建成功后，会生成一个ID，是二层连接对应的唯一标识。 - 约束限制：带“-”的UUID格式。 - 取值范围：不涉及。 - 默认取值：不涉及。
	Id *string `json:"id,omitempty"`

	// - 参数解释：二层连接的名称。 - 约束限制：   - 长度范围为1~64个字符。   - 名称由中文、英文字母、数字、下划线（_）、中划线（-）、点（.）组成。 - 取值范围：不涉及。 - 默认取值：不涉及。
	Name *string `json:"name,omitempty"`
}

func (o ListConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ListConnectionsRequest", string(data)}, " ")
}
