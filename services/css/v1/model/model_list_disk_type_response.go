package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDiskTypeResponse Response Object
type ListDiskTypeResponse struct {

	// **参数解释**： 磁盘类型列表。 **取值范围**： 不涉及
	DiskTypes      *[]DiskType `json:"diskTypes,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListDiskTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDiskTypeResponse struct{}"
	}

	return strings.Join([]string{"ListDiskTypeResponse", string(data)}, " ")
}
