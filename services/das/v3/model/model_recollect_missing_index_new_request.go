package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecollectMissingIndexNewRequest Request Object
type RecollectMissingIndexNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o RecollectMissingIndexNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecollectMissingIndexNewRequest struct{}"
	}

	return strings.Join([]string{"RecollectMissingIndexNewRequest", string(data)}, " ")
}
