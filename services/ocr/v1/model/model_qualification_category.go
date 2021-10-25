package model

import (
	"encoding/json"

	"strings"
)

type QualificationCategory struct {
	// 诚信考核信息（非必有，依赖对应从业资格证板式）。

	Category *string `json:"category,omitempty"`
	// 初次领证日期（非必有，依赖对应从业资格证板式）

	InitialIssueDate *string `json:"initial_issue_date,omitempty"`
	// 有效起始日期（非必有，依赖对应从业资格证板式）

	IssueDate *string `json:"issue_date,omitempty"`
	// 有效期至

	ExpiryDate *string `json:"expiry_date,omitempty"`
}

func (o QualificationCategory) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QualificationCategory struct{}"
	}

	return strings.Join([]string{"QualificationCategory", string(data)}, " ")
}
