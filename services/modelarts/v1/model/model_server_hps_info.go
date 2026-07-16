package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerHpsInfo 服务器归属的超节点信息。
type ServerHpsInfo struct {

	// **参数解释**：Lite Server超节点实例id。 **取值范围**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。
	Id *string `json:"id,omitempty"`

	// **参数解释**：Lite Server超节点名称。 **取值范围**：^[-_.a-zA-Z0-9]{1,64}$。
	Name *string `json:"name,omitempty"`
}

func (o ServerHpsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerHpsInfo struct{}"
	}

	return strings.Join([]string{"ServerHpsInfo", string(data)}, " ")
}
