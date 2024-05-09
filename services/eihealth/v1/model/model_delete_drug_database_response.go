package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDrugDatabaseResponse Response Object
type DeleteDrugDatabaseResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDrugDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDrugDatabaseResponse struct{}"
	}

	return strings.Join([]string{"DeleteDrugDatabaseResponse", string(data)}, " ")
}
