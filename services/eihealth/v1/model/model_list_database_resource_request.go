package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseResourceRequest Request Object
type ListDatabaseResourceRequest struct {
}

func (o ListDatabaseResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseResourceRequest struct{}"
	}

	return strings.Join([]string{"ListDatabaseResourceRequest", string(data)}, " ")
}
