package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IdArray id列表
type IdArray struct {

	// id列表
	Ids []string `json:"ids"`
}

func (o IdArray) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdArray struct{}"
	}

	return strings.Join([]string{"IdArray", string(data)}, " ")
}
