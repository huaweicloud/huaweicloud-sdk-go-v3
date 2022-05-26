package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dbss/v1/model"
)

type SwitchAgentInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchAgentInvoker) Invoke() (*model.SwitchAgentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchAgentResponse), nil
	}
}

type SwitchRiskRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchRiskRuleInvoker) Invoke() (*model.SwitchRiskRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchRiskRuleResponse), nil
	}
}
