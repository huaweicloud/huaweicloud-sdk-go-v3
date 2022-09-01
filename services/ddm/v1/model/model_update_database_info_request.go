package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateDatabaseInfoRequest struct {

	// DDM实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`
}

func (o UpdateDatabaseInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDatabaseInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdateDatabaseInfoRequest", string(data)}, " ")
}
