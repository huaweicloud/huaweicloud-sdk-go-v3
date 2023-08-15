package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateIncidentProcess struct {

	// The name, display only
	ProcessName *string `json:"process_name,omitempty"`

	// The name, display only
	ProcessPath *string `json:"process_path,omitempty"`

	// Id value
	ProcessPid *string `json:"process_pid,omitempty"`

	// Id value
	ProcessUid *string `json:"process_uid,omitempty"`

	// The name, display only
	ProcessCmdline *string `json:"process_cmdline,omitempty"`
}

func (o CreateIncidentProcess) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIncidentProcess struct{}"
	}

	return strings.Join([]string{"CreateIncidentProcess", string(data)}, " ")
}
