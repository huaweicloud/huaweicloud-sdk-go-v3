/*
 * ProjectMan
 *
 * devcloud projectman api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type CheckProjectNameRequestV4 struct {
	// 项目名
	ProjectName string `json:"project_name"`
}

func (o CheckProjectNameRequestV4) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CheckProjectNameRequestV4", string(data)}, " ")
}
