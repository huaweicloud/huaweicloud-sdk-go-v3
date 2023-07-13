package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportDatabaseDataResponse Response Object
type ImportDatabaseDataResponse struct {

	// 数据作业ID
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ImportDatabaseDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportDatabaseDataResponse struct{}"
	}

	return strings.Join([]string{"ImportDatabaseDataResponse", string(data)}, " ")
}
