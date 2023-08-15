package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteComputingResourceRequest Request Object
type DeleteComputingResourceRequest struct {

	// 云服务器ID
	Id string `json:"id"`
}

func (o DeleteComputingResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteComputingResourceRequest struct{}"
	}

	return strings.Join([]string{"DeleteComputingResourceRequest", string(data)}, " ")
}
