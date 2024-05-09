package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDrugDatabaseResponse Response Object
type CreateDrugDatabaseResponse struct {

	// 数据库id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDrugDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDrugDatabaseResponse struct{}"
	}

	return strings.Join([]string{"CreateDrugDatabaseResponse", string(data)}, " ")
}
