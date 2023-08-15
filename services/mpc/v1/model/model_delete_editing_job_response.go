package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEditingJobResponse Response Object
type DeleteEditingJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEditingJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEditingJobResponse struct{}"
	}

	return strings.Join([]string{"DeleteEditingJobResponse", string(data)}, " ")
}
