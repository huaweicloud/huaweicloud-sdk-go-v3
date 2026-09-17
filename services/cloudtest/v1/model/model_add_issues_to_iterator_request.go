package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddIssuesToIteratorRequest Request Object
type AddIssuesToIteratorRequest struct {

	// 项目id
	ProjectUuid string `json:"project_uuid"`

	// 迭代uri
	IteratorUri string `json:"iterator_uri"`

	Body *IssuesInfo `json:"body,omitempty"`
}

func (o AddIssuesToIteratorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddIssuesToIteratorRequest struct{}"
	}

	return strings.Join([]string{"AddIssuesToIteratorRequest", string(data)}, " ")
}
