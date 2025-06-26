package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDiskInfoResponse Response Object
type UpdateDiskInfoResponse struct {

	// 更新磁盘信息成功
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDiskInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDiskInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateDiskInfoResponse", string(data)}, " ")
}
