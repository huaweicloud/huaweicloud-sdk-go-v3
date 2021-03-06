package model

import (
	"encoding/json"

	"strings"
)

//
type JobMetadata struct {
	// 作业的ID。

	Uid *string `json:"uid,omitempty"`
	// 作业的创建时间。

	CreationTimestamp *string `json:"creationTimestamp,omitempty"`
	// 作业的更新时间。

	UpdateTimestamp *string `json:"updateTimestamp,omitempty"`
}

func (o JobMetadata) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "JobMetadata struct{}"
	}

	return strings.Join([]string{"JobMetadata", string(data)}, " ")
}
