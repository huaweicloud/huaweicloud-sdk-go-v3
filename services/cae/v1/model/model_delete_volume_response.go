package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVolumeResponse Response Object
type DeleteVolumeResponse struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion *string `json:"api_version,omitempty"`

	// API类型，固定值“Component”，该值不可修改。
	Kind *string `json:"kind,omitempty"`

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
