package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncAddPersonnelResponse Response Object
type SyncAddPersonnelResponse struct {

	// 添加成功的ID
	AddedIds *[]string `json:"added_ids,omitempty"`

	// 添加失败的ID
	ErrorIds       *[]string `json:"error_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o SyncAddPersonnelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncAddPersonnelResponse struct{}"
	}

	return strings.Join([]string{"SyncAddPersonnelResponse", string(data)}, " ")
}
