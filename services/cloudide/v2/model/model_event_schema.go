package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventSchema the request body of event
type EventSchema struct {

	// the component of the codearts snap
	Component string `json:"component"`

	// the verb of the action
	Verb *string `json:"verb,omitempty"`

	// the the object of the verb
	Object *string `json:"object,omitempty"`

	// the data of the event
	Data *interface{} `json:"data,omitempty"`
}

func (o EventSchema) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventSchema struct{}"
	}

	return strings.Join([]string{"EventSchema", string(data)}, " ")
}
