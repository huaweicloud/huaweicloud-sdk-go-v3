package model

import (
	"encoding/json"

	"strings"
)

//
type SkillResponse struct {
	// 技能ID。

	SkillId string `json:"skill_id"`
	// 技能模型版本。

	SkillVersion string `json:"skill_version"`

	Frame *Frame `json:"frame"`

	Candidate *CandidateIntention `json:"candidate"`
	// 技能是否锁定。

	Locked bool `json:"locked"`
	// 相关意图信息。

	RelatedIntenions []RelatedIntention `json:"related_intenions"`
}

func (o SkillResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SkillResponse struct{}"
	}

	return strings.Join([]string{"SkillResponse", string(data)}, " ")
}
