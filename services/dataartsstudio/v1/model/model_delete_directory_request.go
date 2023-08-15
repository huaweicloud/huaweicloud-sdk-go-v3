package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDirectoryRequest Request Object
type DeleteDirectoryRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 实体id数组
	Ids []int64 `json:"ids"`
}

func (o DeleteDirectoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDirectoryRequest struct{}"
	}

	return strings.Join([]string{"DeleteDirectoryRequest", string(data)}, " ")
}
