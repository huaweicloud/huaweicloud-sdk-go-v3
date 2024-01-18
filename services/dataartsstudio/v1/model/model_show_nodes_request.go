package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNodesRequest Request Object
type ShowNodesRequest struct {

	// 实例id
	Instance string `json:"instance"`

	// 资产guid
	Guid string `json:"guid"`
}

func (o ShowNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNodesRequest struct{}"
	}

	return strings.Join([]string{"ShowNodesRequest", string(data)}, " ")
}
