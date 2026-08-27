package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchBindProvidersRequest Request Object
type BatchBindProvidersRequest struct {

	// 模型组id。
	GroupId string `json:"group_id"`

	Body *BatchBindProvidersReq `json:"body,omitempty"`
}

func (o BatchBindProvidersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchBindProvidersRequest struct{}"
	}

	return strings.Join([]string{"BatchBindProvidersRequest", string(data)}, " ")
}
