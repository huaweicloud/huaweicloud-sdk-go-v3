package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDrugDatabaseFileResponse Response Object
type AddDrugDatabaseFileResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddDrugDatabaseFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDrugDatabaseFileResponse struct{}"
	}

	return strings.Join([]string{"AddDrugDatabaseFileResponse", string(data)}, " ")
}
