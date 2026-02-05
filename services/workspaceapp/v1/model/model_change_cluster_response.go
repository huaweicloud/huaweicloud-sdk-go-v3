package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeClusterResponse Response Object
type ChangeClusterResponse struct {

	// 切换文件夹返回信息。
	Items          *[]ChangeCloudStorageAssignmentInfo `json:"items,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ChangeClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeClusterResponse struct{}"
	}

	return strings.Join([]string{"ChangeClusterResponse", string(data)}, " ")
}
