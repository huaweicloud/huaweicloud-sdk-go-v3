package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ExecuteDeleteimageByIdResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExecuteDeleteimageByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteDeleteimageByIdResponse struct{}"
	}

	return strings.Join([]string{"ExecuteDeleteimageByIdResponse", string(data)}, " ")
}
