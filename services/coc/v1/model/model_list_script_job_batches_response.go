package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScriptJobBatchesResponse Response Object
type ListScriptJobBatchesResponse struct {

	// 展示批次列表
	Data           *[]JobScriptBatchListModel `json:"data,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListScriptJobBatchesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScriptJobBatchesResponse struct{}"
	}

	return strings.Join([]string{"ListScriptJobBatchesResponse", string(data)}, " ")
}
