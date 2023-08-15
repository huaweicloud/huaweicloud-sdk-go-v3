package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteComputingResourceResponse Response Object
type DeleteComputingResourceResponse struct {

	// 被删除计算资源ID。
	ComputingResourceId *string `json:"computing_resource_id,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o DeleteComputingResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteComputingResourceResponse struct{}"
	}

	return strings.Join([]string{"DeleteComputingResourceResponse", string(data)}, " ")
}
