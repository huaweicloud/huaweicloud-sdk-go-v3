package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteResourceRequest struct {

	// 资源id.
	ResourceId string `json:"resource_id"`
}

func (o DeleteResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceRequest struct{}"
	}

	return strings.Join([]string{"DeleteResourceRequest", string(data)}, " ")
}
