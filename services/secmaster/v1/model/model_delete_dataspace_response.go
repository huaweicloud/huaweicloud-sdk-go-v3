package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDataspaceResponse Response Object
type DeleteDataspaceResponse struct {

	// 数据空间ID
	DataspaceId    *string `json:"dataspace_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDataspaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDataspaceResponse struct{}"
	}

	return strings.Join([]string{"DeleteDataspaceResponse", string(data)}, " ")
}
