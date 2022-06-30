package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowStreamCountResponse struct {

	// 采样数据列表。
	DataList *[]StreamCountData `json:"data_list,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowStreamCountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStreamCountResponse struct{}"
	}

	return strings.Join([]string{"ShowStreamCountResponse", string(data)}, " ")
}
