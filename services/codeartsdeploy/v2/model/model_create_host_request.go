package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHostRequest Request Object
type CreateHostRequest struct {

	// 主机集群id
	GroupId string `json:"group_id"`

	Body *CreateHostRequestBody `json:"body,omitempty"`
}

func (o CreateHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHostRequest struct{}"
	}

	return strings.Join([]string{"CreateHostRequest", string(data)}, " ")
}
