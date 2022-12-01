package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListWorkSpacesResponse struct {

	// 返回总条目数
	Count *int32 `json:"count,omitempty"`

	Workspaces     *WorkspaceListElem `json:"workspaces,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListWorkSpacesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkSpacesResponse struct{}"
	}

	return strings.Join([]string{"ListWorkSpacesResponse", string(data)}, " ")
}
