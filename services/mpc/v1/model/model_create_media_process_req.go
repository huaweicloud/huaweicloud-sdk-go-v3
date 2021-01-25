package model

import (
	"encoding/json"

	"strings"
)

type CreateMediaProcessReq struct {
	Input  *ObsObjInfo `json:"input,omitempty"`
	Output *ObsObjInfo `json:"output,omitempty"`
	// 模板ID
	TemplateId *string `json:"template_id,omitempty"`
}

func (o CreateMediaProcessReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMediaProcessReq struct{}"
	}

	return strings.Join([]string{"CreateMediaProcessReq", string(data)}, " ")
}
