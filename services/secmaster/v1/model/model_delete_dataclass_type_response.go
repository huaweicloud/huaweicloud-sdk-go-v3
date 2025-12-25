package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDataclassTypeResponse Response Object
type DeleteDataclassTypeResponse struct {

	// 关联资源信息列表
	Body           *[]DeleteRelationResource `json:"body,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o DeleteDataclassTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDataclassTypeResponse struct{}"
	}

	return strings.Join([]string{"DeleteDataclassTypeResponse", string(data)}, " ")
}
