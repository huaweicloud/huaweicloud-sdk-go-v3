package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseAvailableVersionsRequest Request Object
type ListDatabaseAvailableVersionsRequest struct {

	// DDM实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ListDatabaseAvailableVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseAvailableVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListDatabaseAvailableVersionsRequest", string(data)}, " ")
}
