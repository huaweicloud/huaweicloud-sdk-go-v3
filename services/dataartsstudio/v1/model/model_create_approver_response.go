package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateApproverResponse Response Object
type CreateApproverResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o CreateApproverResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApproverResponse struct{}"
	}

	return strings.Join([]string{"CreateApproverResponse", string(data)}, " ")
}
