package vmock

import (
	"k8s.io/cloud-provider-alibaba-cloud/pkg/provider"
	"k8s.io/cloud-provider-alibaba-cloud/pkg/provider/alibaba"
	"k8s.io/cloud-provider-alibaba-cloud/pkg/provider/alibaba/base"
)

func NewMockCloud(auth *base.ClientMgr) prvd.Provider {

	return &MockCloud{
		IMetaData: auth.Meta,
		MockECS:   NewMockECS(auth),
		MockCLB:   NewMockCLB(auth),
		MockPVTZ:  NewMockPVTZ(auth),
		MockVPC:   NewMockVPC(auth),
	}
}

var _ prvd.Provider = alibaba.AlibabaCloud{}

// MockCloud for unit test
type MockCloud struct {
	*MockECS
	*MockPVTZ
	*MockVPC
	*MockCLB
	prvd.IMetaData
}
