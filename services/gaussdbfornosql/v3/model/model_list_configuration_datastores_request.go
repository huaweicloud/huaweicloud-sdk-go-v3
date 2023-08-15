package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigurationDatastoresRequest Request Object
type ListConfigurationDatastoresRequest struct {
}

func (o ListConfigurationDatastoresRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationDatastoresRequest struct{}"
	}

	return strings.Join([]string{"ListConfigurationDatastoresRequest", string(data)}, " ")
}
