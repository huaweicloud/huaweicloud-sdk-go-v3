package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateUsersRequest struct {

	// DDM实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *CreateUsersReq `json:"body,omitempty" xml:"body"`
}

func (o CreateUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUsersRequest struct{}"
	}

	return strings.Join([]string{"CreateUsersRequest", string(data)}, " ")
}
