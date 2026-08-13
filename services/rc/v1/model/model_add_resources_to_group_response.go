package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddResourcesToGroupResponse Response Object
type AddResourcesToGroupResponse struct {

	// 成功添加的资源列表
	Resources      *[]ResourceEntity `json:"resources,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o AddResourcesToGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddResourcesToGroupResponse struct{}"
	}

	return strings.Join([]string{"AddResourcesToGroupResponse", string(data)}, " ")
}
