package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserConnectionRequest Request Object
type ListUserConnectionRequest struct {

	// 单次查询的大小[1-100]
	Limit int32 `json:"limit"`

	// 查询的偏移量
	Offset int32 `json:"offset"`

	Body *ListUserConnectionReq `json:"body,omitempty"`
}

func (o ListUserConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserConnectionRequest struct{}"
	}

	return strings.Join([]string{"ListUserConnectionRequest", string(data)}, " ")
}
