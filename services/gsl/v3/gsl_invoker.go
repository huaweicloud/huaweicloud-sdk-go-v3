package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/gsl/v3/model"
)

type ListProPricePlansInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProPricePlansInvoker) Invoke() (*model.ListProPricePlansResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProPricePlansResponse), nil
	}
}

type DeleteRealNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRealNameInvoker) Invoke() (*model.DeleteRealNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRealNameResponse), nil
	}
}

type EnableSimCardInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableSimCardInvoker) Invoke() (*model.EnableSimCardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableSimCardResponse), nil
	}
}

type ListSimCardsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSimCardsInvoker) Invoke() (*model.ListSimCardsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSimCardsResponse), nil
	}
}

type RegisterImeiInvoker struct {
	*invoker.BaseInvoker
}

func (i *RegisterImeiInvoker) Invoke() (*model.RegisterImeiResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterImeiResponse), nil
	}
}

type ResetSimCardInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetSimCardInvoker) Invoke() (*model.ResetSimCardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetSimCardResponse), nil
	}
}

type SetExceedCutNetInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetExceedCutNetInvoker) Invoke() (*model.SetExceedCutNetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetExceedCutNetResponse), nil
	}
}

type SetSpeedValueInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetSpeedValueInvoker) Invoke() (*model.SetSpeedValueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetSpeedValueResponse), nil
	}
}

type ShowRealNamedInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRealNamedInvoker) Invoke() (*model.ShowRealNamedResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRealNamedResponse), nil
	}
}

type ShowSimCardInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSimCardInvoker) Invoke() (*model.ShowSimCardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSimCardResponse), nil
	}
}

type StartStopNetInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartStopNetInvoker) Invoke() (*model.StartStopNetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartStopNetResponse), nil
	}
}

type StopSimCardInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopSimCardInvoker) Invoke() (*model.StopSimCardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopSimCardResponse), nil
	}
}

type ListSimPoolMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSimPoolMembersInvoker) Invoke() (*model.ListSimPoolMembersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSimPoolMembersResponse), nil
	}
}

type ListSimPoolsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSimPoolsInvoker) Invoke() (*model.ListSimPoolsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSimPoolsResponse), nil
	}
}

type ListFlowBySimCardsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlowBySimCardsInvoker) Invoke() (*model.ListFlowBySimCardsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlowBySimCardsResponse), nil
	}
}

type ListSimPricePlansInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSimPricePlansInvoker) Invoke() (*model.ListSimPricePlansResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSimPricePlansResponse), nil
	}
}
