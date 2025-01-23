package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIpGroupRelatedListenersResponse Response Object
type ShowIpGroupRelatedListenersResponse struct {

	// IP地址组关联的所有监听器ID列表
	Listeners      *[]ListenerRef `json:"listeners,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowIpGroupRelatedListenersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpGroupRelatedListenersResponse struct{}"
	}

	return strings.Join([]string{"ShowIpGroupRelatedListenersResponse", string(data)}, " ")
}
