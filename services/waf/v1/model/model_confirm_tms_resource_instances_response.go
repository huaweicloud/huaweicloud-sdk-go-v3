package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmTmsResourceInstancesResponse Response Object
type ConfirmTmsResourceInstancesResponse struct {
	TotalCount *int32 `json:"total_count,omitempty"`

	Resources      *[]TmsResourceInstance `json:"resources,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ConfirmTmsResourceInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmTmsResourceInstancesResponse struct{}"
	}

	return strings.Join([]string{"ConfirmTmsResourceInstancesResponse", string(data)}, " ")
}
