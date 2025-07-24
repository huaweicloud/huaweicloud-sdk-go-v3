package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListShareTypesResponse Response Object
type ListShareTypesResponse struct {

	// 文件系统类型和配额列表
	ShareTypes *[]ShareTypeResponseBody `json:"share_types,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListShareTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListShareTypesResponse struct{}"
	}

	return strings.Join([]string{"ListShareTypesResponse", string(data)}, " ")
}
