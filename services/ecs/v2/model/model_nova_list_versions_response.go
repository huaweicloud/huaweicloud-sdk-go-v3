package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NovaListVersionsResponse Response Object
type NovaListVersionsResponse struct {

	// API版本信息列表
	Versions       *[]NovaVersion `json:"versions,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o NovaListVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaListVersionsResponse struct{}"
	}

	return strings.Join([]string{"NovaListVersionsResponse", string(data)}, " ")
}
