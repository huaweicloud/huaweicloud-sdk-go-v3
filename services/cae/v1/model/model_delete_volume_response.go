package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVolumeResponse Response Object
type DeleteVolumeResponse struct {
	ApiVersion *ApiVersionObj `json:"api_version,omitempty"`

	Kind *ComponentKindObj `json:"kind,omitempty"`

	// 挂载组件列表。
	Items          *[]MountComponent `json:"items,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DeleteVolumeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVolumeResponse struct{}"
	}

	return strings.Join([]string{"DeleteVolumeResponse", string(data)}, " ")
}
