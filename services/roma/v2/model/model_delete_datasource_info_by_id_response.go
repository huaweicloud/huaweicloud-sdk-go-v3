package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDatasourceInfoByIdResponse Response Object
type DeleteDatasourceInfoByIdResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDatasourceInfoByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatasourceInfoByIdResponse struct{}"
	}

	return strings.Join([]string{"DeleteDatasourceInfoByIdResponse", string(data)}, " ")
}
