package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBackupsRequest Request Object
type ListBackupsRequest struct {

	// 每页资源数量的最大值，默认为10。
	Limit *int32 `json:"limit,omitempty"`

	// 本次请求的起始位置，默认为0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListBackupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupsRequest struct{}"
	}

	return strings.Join([]string{"ListBackupsRequest", string(data)}, " ")
}
