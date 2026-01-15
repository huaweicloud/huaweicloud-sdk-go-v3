package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteColumnInfosResponse Response Object
type DeleteColumnInfosResponse struct {

	// 删除结果
	DeleteResult   *bool `json:"delete_result,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o DeleteColumnInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteColumnInfosResponse struct{}"
	}

	return strings.Join([]string{"DeleteColumnInfosResponse", string(data)}, " ")
}
