package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDataobjectsResponse Response Object
type DeleteDataobjectsResponse struct {

	// 失败id
	ErrorIds *[]string `json:"error_ids,omitempty"`

	// 成功id
	SuccessIds     *[]string `json:"success_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o DeleteDataobjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDataobjectsResponse struct{}"
	}

	return strings.Join([]string{"DeleteDataobjectsResponse", string(data)}, " ")
}
