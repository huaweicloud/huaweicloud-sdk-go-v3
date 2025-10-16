package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateProtectStatusResponse Response Object
type BatchUpdateProtectStatusResponse struct {

	// 云模式防护域名的数量
	Total *int32 `json:"total,omitempty"`

	// 详细的云模式防护域名列表信息
	Items          *[]CloudWafHostItem `json:"items,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o BatchUpdateProtectStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateProtectStatusResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateProtectStatusResponse", string(data)}, " ")
}
