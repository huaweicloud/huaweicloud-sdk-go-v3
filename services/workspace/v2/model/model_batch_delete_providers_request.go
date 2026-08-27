package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteProvidersRequest Request Object
type BatchDeleteProvidersRequest struct {

	// 模型组id。
	GroupId string `json:"group_id"`

	Body *BatchDeleteProvidersReq `json:"body,omitempty"`
}

func (o BatchDeleteProvidersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteProvidersRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteProvidersRequest", string(data)}, " ")
}
