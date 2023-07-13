package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteCorpAdminsResponse Response Object
type BatchDeleteCorpAdminsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteCorpAdminsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteCorpAdminsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteCorpAdminsResponse", string(data)}, " ")
}
