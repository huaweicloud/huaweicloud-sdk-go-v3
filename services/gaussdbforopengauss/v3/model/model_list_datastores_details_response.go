package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatastoresDetailsResponse Response Object
type ListDatastoresDetailsResponse struct {
	Datastores     *[]DatastoresResult `json:"datastores,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListDatastoresDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatastoresDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListDatastoresDetailsResponse", string(data)}, " ")
}
