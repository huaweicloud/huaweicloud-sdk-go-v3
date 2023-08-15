package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseUsersRequest Request Object
type ListDatabaseUsersRequest struct {

	// cluster_id
	ClusterId string `json:"cluster_id"`
}

func (o ListDatabaseUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseUsersRequest struct{}"
	}

	return strings.Join([]string{"ListDatabaseUsersRequest", string(data)}, " ")
}
