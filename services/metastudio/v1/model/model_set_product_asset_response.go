package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetProductAssetResponse Response Object
type SetProductAssetResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetProductAssetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetProductAssetResponse struct{}"
	}

	return strings.Join([]string{"SetProductAssetResponse", string(data)}, " ")
}
