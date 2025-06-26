package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceRegistryResponse Response Object
type CreateInstanceRegistryResponse struct {

	// 命名空间ID
	Id             *int32 `json:"id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateInstanceRegistryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceRegistryResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceRegistryResponse", string(data)}, " ")
}
