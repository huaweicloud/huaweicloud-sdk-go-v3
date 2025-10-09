package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IdListWrapper struct {

	// 监听器ID。
	Id *string `json:"id,omitempty"`
}

func (o IdListWrapper) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdListWrapper struct{}"
	}

	return strings.Join([]string{"IdListWrapper", string(data)}, " ")
}
