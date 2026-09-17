package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompareSlowLogTemplatesRequest Request Object
type CompareSlowLogTemplatesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *CompareSlowLogTemplatesRequestBody `json:"body,omitempty"`
}

func (o CompareSlowLogTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareSlowLogTemplatesRequest struct{}"
	}

	return strings.Join([]string{"CompareSlowLogTemplatesRequest", string(data)}, " ")
}
