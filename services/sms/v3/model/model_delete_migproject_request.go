package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteMigprojectRequest struct {
	// 需要删除的迁移项目的id

	MigProjectId string `json:"mig_project_id"`
}

func (o DeleteMigprojectRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteMigprojectRequest struct{}"
	}

	return strings.Join([]string{"DeleteMigprojectRequest", string(data)}, " ")
}
