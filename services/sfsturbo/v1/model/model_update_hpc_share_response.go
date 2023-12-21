package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHpcShareResponse Response Object
type UpdateHpcShareResponse struct {

	// 文件系统冷数据淘汰时间
	GcTime *int32 `json:"gc_time,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateHpcShareResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHpcShareResponse struct{}"
	}

	return strings.Join([]string{"UpdateHpcShareResponse", string(data)}, " ")
}
