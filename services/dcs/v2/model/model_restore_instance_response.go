package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RestoreInstanceResponse struct {
	// 恢复记录ID。

	RestoreId      *string `json:"restore_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestoreInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreInstanceResponse struct{}"
	}

	return strings.Join([]string{"RestoreInstanceResponse", string(data)}, " ")
}
