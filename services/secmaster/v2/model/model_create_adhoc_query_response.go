package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAdhocQueryResponse Response Object
type CreateAdhocQueryResponse struct {

	// 操作ID
	OperateId      *string `json:"operate_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAdhocQueryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAdhocQueryResponse struct{}"
	}

	return strings.Join([]string{"CreateAdhocQueryResponse", string(data)}, " ")
}
