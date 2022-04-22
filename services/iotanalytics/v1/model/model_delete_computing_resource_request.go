package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteComputingResourceRequest struct {

	// 计算资源ID。
	ComputingResourceId string `json:"computing_resource_id"`
}

func (o DeleteComputingResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteComputingResourceRequest struct{}"
	}

	return strings.Join([]string{"DeleteComputingResourceRequest", string(data)}, " ")
}
