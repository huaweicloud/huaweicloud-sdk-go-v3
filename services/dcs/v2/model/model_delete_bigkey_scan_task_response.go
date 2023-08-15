package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBigkeyScanTaskResponse Response Object
type DeleteBigkeyScanTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteBigkeyScanTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBigkeyScanTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteBigkeyScanTaskResponse", string(data)}, " ")
}
