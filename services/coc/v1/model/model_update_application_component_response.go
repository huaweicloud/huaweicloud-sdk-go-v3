package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateApplicationComponentResponse Response Object
type UpdateApplicationComponentResponse struct {

	// 修改的组件id。
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateApplicationComponentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApplicationComponentResponse struct{}"
	}

	return strings.Join([]string{"UpdateApplicationComponentResponse", string(data)}, " ")
}
