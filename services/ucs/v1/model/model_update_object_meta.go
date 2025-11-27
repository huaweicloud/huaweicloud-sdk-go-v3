package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateObjectMeta Object Meta Data
type UpdateObjectMeta struct {

	// 集群标注
	Annotations *interface{} `json:"annotations,omitempty"`
}

func (o UpdateObjectMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateObjectMeta struct{}"
	}

	return strings.Join([]string{"UpdateObjectMeta", string(data)}, " ")
}
