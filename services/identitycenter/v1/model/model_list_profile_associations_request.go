package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProfileAssociationsRequest Request Object
type ListProfileAssociationsRequest struct {
	InstanceId string `json:"instance_id"`

	// 待查询的profile唯一标识
	ProfileId string `json:"profile_id"`
}

func (o ListProfileAssociationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProfileAssociationsRequest struct{}"
	}

	return strings.Join([]string{"ListProfileAssociationsRequest", string(data)}, " ")
}
