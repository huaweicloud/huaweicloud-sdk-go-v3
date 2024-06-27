package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteTestCasesResponse Response Object
type BatchDeleteTestCasesResponse struct {
	Value          *string `json:"value,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeleteTestCasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTestCasesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteTestCasesResponse", string(data)}, " ")
}
