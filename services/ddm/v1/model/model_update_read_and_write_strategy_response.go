package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateReadAndWriteStrategyResponse struct {

	// 操作是否成功。
	Success *bool `json:"success,omitempty" xml:"success"`

	// DDM实例ID。
	InstanceId     *string `json:"instance_id,omitempty" xml:"instance_id"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateReadAndWriteStrategyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateReadAndWriteStrategyResponse struct{}"
	}

	return strings.Join([]string{"UpdateReadAndWriteStrategyResponse", string(data)}, " ")
}
