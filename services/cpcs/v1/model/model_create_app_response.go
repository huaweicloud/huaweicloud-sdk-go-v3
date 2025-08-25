package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppResponse Response Object
type CreateAppResponse struct {

	// 创建出来的应用的ID
	AppId          *string `json:"app_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppResponse struct{}"
	}

	return strings.Join([]string{"CreateAppResponse", string(data)}, " ")
}
