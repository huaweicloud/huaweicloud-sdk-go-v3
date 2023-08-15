package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecordResponse Response Object
type RecordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordResponse struct{}"
	}

	return strings.Join([]string{"RecordResponse", string(data)}, " ")
}
