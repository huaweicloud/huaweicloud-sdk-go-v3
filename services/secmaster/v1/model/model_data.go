package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Data 关联资源名称/id
type Data struct {

	// 关联项名称
	Name *string `json:"name,omitempty"`

	// 关联项id
	Id *string `json:"id,omitempty"`
}

func (o Data) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Data struct{}"
	}

	return strings.Join([]string{"Data", string(data)}, " ")
}
