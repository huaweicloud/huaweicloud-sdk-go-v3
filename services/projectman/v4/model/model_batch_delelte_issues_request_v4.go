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

type BatchDelelteIssuesRequestV4 struct {
	// 工作项的id
	IssueIds []int32 `json:"issue_ids"`
}

func (o BatchDelelteIssuesRequestV4) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchDelelteIssuesRequestV4", string(data)}, " ")
}
