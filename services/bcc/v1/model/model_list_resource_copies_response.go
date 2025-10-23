package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceCopiesResponse Response Object
type ListResourceCopiesResponse struct {

	// 副本总数
	Count *int64 `json:"count,omitempty"`

	// 副本列表
	Copies *[]ResourceCopyItem `json:"copies,omitempty"`

	// 下一页的marker
	NextMarker     *string `json:"next_marker,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListResourceCopiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceCopiesResponse struct{}"
	}

	return strings.Join([]string{"ListResourceCopiesResponse", string(data)}, " ")
}
