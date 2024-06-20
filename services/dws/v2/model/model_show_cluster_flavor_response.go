package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterFlavorResponse Response Object
type ShowClusterFlavorResponse struct {
	Flavor         *ClusterFlavorDetailInfo `json:"flavor,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowClusterFlavorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterFlavorResponse struct{}"
	}

	return strings.Join([]string{"ShowClusterFlavorResponse", string(data)}, " ")
}
