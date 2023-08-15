package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIncident 删除事件
type DeleteIncident struct {

	// 删除事件的ID列表
	Ids *[]string `json:"ids,omitempty"`
}

func (o DeleteIncident) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIncident struct{}"
	}

	return strings.Join([]string{"DeleteIncident", string(data)}, " ")
}
