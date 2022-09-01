package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type MigrateNodeResponse struct {

	// API版本，固定值“v3”。
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion"`

	// API类型，固定值“MigrateNodesTask”。
	Kind *string `json:"kind,omitempty" xml:"kind"`

	Spec *MigrateNodesSpec `json:"spec,omitempty" xml:"spec"`

	Status         *TaskStatus `json:"status,omitempty" xml:"status"`
	HttpStatusCode int         `json:"-"`
}

func (o MigrateNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateNodeResponse struct{}"
	}

	return strings.Join([]string{"MigrateNodeResponse", string(data)}, " ")
}
