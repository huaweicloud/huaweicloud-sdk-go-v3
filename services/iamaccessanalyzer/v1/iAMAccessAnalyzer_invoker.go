package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iamaccessanalyzer/v1/model"
)

type CreateAnalyzerInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAnalyzerInvoker) Invoke() (*model.CreateAnalyzerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAnalyzerResponse), nil
	}
}

type DeleteAnalyzerInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAnalyzerInvoker) Invoke() (*model.DeleteAnalyzerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAnalyzerResponse), nil
	}
}

type ListAnalyzersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAnalyzersInvoker) Invoke() (*model.ListAnalyzersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAnalyzersResponse), nil
	}
}

type ShowAnalyzerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAnalyzerInvoker) Invoke() (*model.ShowAnalyzerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAnalyzerResponse), nil
	}
}

type StartResourceScanInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartResourceScanInvoker) Invoke() (*model.StartResourceScanResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartResourceScanResponse), nil
	}
}

type ApplyArchiveRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ApplyArchiveRuleInvoker) Invoke() (*model.ApplyArchiveRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApplyArchiveRuleResponse), nil
	}
}

type CreateArchiveRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateArchiveRuleInvoker) Invoke() (*model.CreateArchiveRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateArchiveRuleResponse), nil
	}
}

type DeleteArchiveRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteArchiveRuleInvoker) Invoke() (*model.DeleteArchiveRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteArchiveRuleResponse), nil
	}
}

type ListArchiveRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListArchiveRulesInvoker) Invoke() (*model.ListArchiveRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListArchiveRulesResponse), nil
	}
}

type ShowArchiveRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowArchiveRuleInvoker) Invoke() (*model.ShowArchiveRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowArchiveRuleResponse), nil
	}
}

type UpdateArchiveRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateArchiveRuleInvoker) Invoke() (*model.UpdateArchiveRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateArchiveRuleResponse), nil
	}
}

type ListFindingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFindingsInvoker) Invoke() (*model.ListFindingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFindingsResponse), nil
	}
}

type ShowFindingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFindingInvoker) Invoke() (*model.ShowFindingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFindingResponse), nil
	}
}

type UpdateFindingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFindingsInvoker) Invoke() (*model.UpdateFindingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFindingsResponse), nil
	}
}

type CreateAccessPreviewInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAccessPreviewInvoker) Invoke() (*model.CreateAccessPreviewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAccessPreviewResponse), nil
	}
}

type ListAccessPreviewFindingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccessPreviewFindingsInvoker) Invoke() (*model.ListAccessPreviewFindingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccessPreviewFindingsResponse), nil
	}
}

type ListAccessPreviewsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccessPreviewsInvoker) Invoke() (*model.ListAccessPreviewsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccessPreviewsResponse), nil
	}
}

type ShowAccessPreviewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAccessPreviewInvoker) Invoke() (*model.ShowAccessPreviewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAccessPreviewResponse), nil
	}
}

type TagResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *TagResourceInvoker) Invoke() (*model.TagResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.TagResourceResponse), nil
	}
}

type UntagResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UntagResourceInvoker) Invoke() (*model.UntagResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UntagResourceResponse), nil
	}
}

type CheckNoNewAccessInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckNoNewAccessInvoker) Invoke() (*model.CheckNoNewAccessResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckNoNewAccessResponse), nil
	}
}

type ValidatePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ValidatePolicyInvoker) Invoke() (*model.ValidatePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ValidatePolicyResponse), nil
	}
}
