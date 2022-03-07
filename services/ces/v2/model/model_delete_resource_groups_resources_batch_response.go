package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteResourceGroupsResourcesBatchResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteResourceGroupsResourcesBatchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceGroupsResourcesBatchResponse struct{}"
	}

	return strings.Join([]string{"DeleteResourceGroupsResourcesBatchResponse", string(data)}, " ")
}
