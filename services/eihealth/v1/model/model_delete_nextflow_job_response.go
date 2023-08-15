package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNextflowJobResponse Response Object
type DeleteNextflowJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteNextflowJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNextflowJobResponse struct{}"
	}

	return strings.Join([]string{"DeleteNextflowJobResponse", string(data)}, " ")
}
