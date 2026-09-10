package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachDevServerPortResponse Response Object
type DetachDevServerPortResponse struct {

	// **参数解释**：卸载的网卡ID。 **取值范围**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。
	PortId         *string `json:"port_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DetachDevServerPortResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachDevServerPortResponse struct{}"
	}

	return strings.Join([]string{"DetachDevServerPortResponse", string(data)}, " ")
}
