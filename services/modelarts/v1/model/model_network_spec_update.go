package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NetworkSpecUpdate 网络资源描述更新信息。
type NetworkSpecUpdate struct {
	Connection *NetworkConnection `json:"connection,omitempty"`
}

func (o NetworkSpecUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkSpecUpdate struct{}"
	}

	return strings.Join([]string{"NetworkSpecUpdate", string(data)}, " ")
}
