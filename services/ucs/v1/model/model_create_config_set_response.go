package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConfigSetResponse Response Object
type CreateConfigSetResponse struct {

	// 配置集合ID
	Uid            *string `json:"uid,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateConfigSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigSetResponse struct{}"
	}

	return strings.Join([]string{"CreateConfigSetResponse", string(data)}, " ")
}
