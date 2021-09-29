package model

import (
	"encoding/json"

	"strings"
)

type LogicController struct {
	// for_loop_params

	ForLoopParams *string `json:"for_loop_params,omitempty"`
	// condition

	Condition *string `json:"condition,omitempty"`
}

func (o LogicController) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "LogicController struct{}"
	}

	return strings.Join([]string{"LogicController", string(data)}, " ")
}
