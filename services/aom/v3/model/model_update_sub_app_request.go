package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSubAppRequest Request Object
type UpdateSubAppRequest struct {

	// 子应用id
	SubAppId string `json:"sub_app_id"`

	Body *SubAppUpdateParam `json:"body,omitempty"`
}

func (o UpdateSubAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubAppRequest struct{}"
	}

	return strings.Join([]string{"UpdateSubAppRequest", string(data)}, " ")
}
