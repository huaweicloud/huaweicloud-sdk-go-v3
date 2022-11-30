package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListGaussDbDatastoresResponse struct {
	Datastores     *[]DatastoresResult `json:"datastores,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListGaussDbDatastoresResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGaussDbDatastoresResponse struct{}"
	}

	return strings.Join([]string{"ListGaussDbDatastoresResponse", string(data)}, " ")
}
