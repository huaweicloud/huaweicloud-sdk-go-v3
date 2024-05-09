package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDrugDatabaseResponse Response Object
type UpdateDrugDatabaseResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDrugDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDrugDatabaseResponse struct{}"
	}

	return strings.Join([]string{"UpdateDrugDatabaseResponse", string(data)}, " ")
}
