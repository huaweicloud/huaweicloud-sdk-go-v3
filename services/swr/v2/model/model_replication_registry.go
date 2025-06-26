package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReplicationRegistry struct {

	// The registry ID.
	Id int64 `json:"id"`
}

func (o ReplicationRegistry) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplicationRegistry struct{}"
	}

	return strings.Join([]string{"ReplicationRegistry", string(data)}, " ")
}
