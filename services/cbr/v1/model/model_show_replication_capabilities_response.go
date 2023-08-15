package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReplicationCapabilitiesResponse Response Object
type ShowReplicationCapabilitiesResponse struct {

	// 支持复制的区域列表
	Regions        *[]ProtectableReplicationCapabilitiesRespRegion `json:"regions,omitempty"`
	HttpStatusCode int                                             `json:"-"`
}

func (o ShowReplicationCapabilitiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReplicationCapabilitiesResponse struct{}"
	}

	return strings.Join([]string{"ShowReplicationCapabilitiesResponse", string(data)}, " ")
}
