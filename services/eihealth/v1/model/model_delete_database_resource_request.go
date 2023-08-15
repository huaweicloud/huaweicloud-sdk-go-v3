package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDatabaseResourceRequest Request Object
type DeleteDatabaseResourceRequest struct {

	// 数据库ID
	Id string `json:"id"`
}

func (o DeleteDatabaseResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabaseResourceRequest struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseResourceRequest", string(data)}, " ")
}
