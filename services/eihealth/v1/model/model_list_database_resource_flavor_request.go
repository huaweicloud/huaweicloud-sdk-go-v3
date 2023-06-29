package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseResourceFlavorRequest Request Object
type ListDatabaseResourceFlavorRequest struct {
}

func (o ListDatabaseResourceFlavorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseResourceFlavorRequest struct{}"
	}

	return strings.Join([]string{"ListDatabaseResourceFlavorRequest", string(data)}, " ")
}
