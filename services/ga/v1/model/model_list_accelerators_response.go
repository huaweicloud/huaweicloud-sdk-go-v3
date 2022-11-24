package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAcceleratorsResponse struct {

	// 全球加速器列表。
	Accelerators *[]AcceleratorDetail `json:"accelerators,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListAcceleratorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAcceleratorsResponse struct{}"
	}

	return strings.Join([]string{"ListAcceleratorsResponse", string(data)}, " ")
}
