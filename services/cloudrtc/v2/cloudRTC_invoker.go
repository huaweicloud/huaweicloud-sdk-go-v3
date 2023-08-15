package v2

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/cloudrtc/v2/model"
)

type CreateAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAppInvoker) Invoke() (*model.CreateAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAppResponse), nil
	}
}

type CreateIndividualStreamJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateIndividualStreamJobInvoker) Invoke() (*model.CreateIndividualStreamJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateIndividualStreamJobResponse), nil
	}
}

type CreateMixJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMixJobInvoker) Invoke() (*model.CreateMixJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMixJobResponse), nil
	}
}

type CreateRecordRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRecordRuleInvoker) Invoke() (*model.CreateRecordRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRecordRuleResponse), nil
	}
}

type DeleteAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAppInvoker) Invoke() (*model.DeleteAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAppResponse), nil
	}
}

type DeleteRecordRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRecordRuleInvoker) Invoke() (*model.DeleteRecordRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRecordRuleResponse), nil
	}
}

type ListAppsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppsInvoker) Invoke() (*model.ListAppsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppsResponse), nil
	}
}

type ListRecordRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRecordRulesInvoker) Invoke() (*model.ListRecordRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRecordRulesResponse), nil
	}
}

type RemoveRoomInvoker struct {
	*invoker.BaseInvoker
}

func (i *RemoveRoomInvoker) Invoke() (*model.RemoveRoomResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RemoveRoomResponse), nil
	}
}

type RemoveUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *RemoveUsersInvoker) Invoke() (*model.RemoveUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RemoveUsersResponse), nil
	}
}

type ShowAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppInvoker) Invoke() (*model.ShowAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppResponse), nil
	}
}

type ShowAutoRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAutoRecordInvoker) Invoke() (*model.ShowAutoRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAutoRecordResponse), nil
	}
}

type ShowIndividualStreamJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIndividualStreamJobInvoker) Invoke() (*model.ShowIndividualStreamJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIndividualStreamJobResponse), nil
	}
}

type ShowMixJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMixJobInvoker) Invoke() (*model.ShowMixJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMixJobResponse), nil
	}
}

type ShowRecordCallbackInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRecordCallbackInvoker) Invoke() (*model.ShowRecordCallbackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRecordCallbackResponse), nil
	}
}

type ShowRecordRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRecordRuleInvoker) Invoke() (*model.ShowRecordRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRecordRuleResponse), nil
	}
}

type ShowUrlAuthInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUrlAuthInvoker) Invoke() (*model.ShowUrlAuthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUrlAuthResponse), nil
	}
}

type StartAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartAppInvoker) Invoke() (*model.StartAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartAppResponse), nil
	}
}

type StopAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopAppInvoker) Invoke() (*model.StopAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopAppResponse), nil
	}
}

type StopIndividualStreamJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopIndividualStreamJobInvoker) Invoke() (*model.StopIndividualStreamJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopIndividualStreamJobResponse), nil
	}
}

type StopMixJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopMixJobInvoker) Invoke() (*model.StopMixJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopMixJobResponse), nil
	}
}

type UpdateAutoRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAutoRecordInvoker) Invoke() (*model.UpdateAutoRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAutoRecordResponse), nil
	}
}

type UpdateIndividualStreamJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateIndividualStreamJobInvoker) Invoke() (*model.UpdateIndividualStreamJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateIndividualStreamJobResponse), nil
	}
}

type UpdateMixJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMixJobInvoker) Invoke() (*model.UpdateMixJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMixJobResponse), nil
	}
}

type UpdateRecordCallbackInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRecordCallbackInvoker) Invoke() (*model.UpdateRecordCallbackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRecordCallbackResponse), nil
	}
}

type UpdateRecordRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRecordRuleInvoker) Invoke() (*model.UpdateRecordRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRecordRuleResponse), nil
	}
}

type UpdateUrlAuthInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateUrlAuthInvoker) Invoke() (*model.UpdateUrlAuthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateUrlAuthResponse), nil
	}
}
