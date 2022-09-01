package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListServersByTagResponse struct {

	// 返回的云服务器列表。
	Resources *[]ServerResource `json:"resources,omitempty" xml:"resources"`

	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty" xml:"total_count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListServersByTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServersByTagResponse struct{}"
	}

	return strings.Join([]string{"ListServersByTagResponse", string(data)}, " ")
}
