package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMemberRequest Request Object
type CreateMemberRequest struct {

	// 后端云服务器组id
	PoolId string `json:"pool_id"`

	Body *CreateMemberRequestBody `json:"body,omitempty"`
}

func (o CreateMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMemberRequest struct{}"
	}

	return strings.Join([]string{"CreateMemberRequest", string(data)}, " ")
}
