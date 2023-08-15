package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBackup2Request Request Object
type CreateBackup2Request struct {

	// 图ID。
	GraphId string `json:"graph_id"`
}

func (o CreateBackup2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBackup2Request struct{}"
	}

	return strings.Join([]string{"CreateBackup2Request", string(data)}, " ")
}
