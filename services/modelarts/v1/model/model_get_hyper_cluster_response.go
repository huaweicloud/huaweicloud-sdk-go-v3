package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetHyperClusterResponse Response Object
type GetHyperClusterResponse struct {

	// **参数解释**：hyper cluster的ID。 **取值范围**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。
	Id *string `json:"id,omitempty"`

	// **参数解释**：hyper cluster的名称。 **取值范围**：^[-_.a-zA-Z0-9]{1,64}$。
	Name *string `json:"name,omitempty"`

	// **参数解释**：网络信息。
	NetworkInfo    *[]HyperClusterNetworkInfo `json:"network_info,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o GetHyperClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetHyperClusterResponse struct{}"
	}

	return strings.Join([]string{"GetHyperClusterResponse", string(data)}, " ")
}
