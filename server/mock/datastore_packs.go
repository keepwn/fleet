// Automatically generated by mockimpl. DO NOT EDIT!

package mock

import "github.com/fleetdm/fleet/v4/server/fleet"

var _ fleet.PackStore = (*PackStore)(nil)

type ApplyPackSpecsFunc func(specs []*fleet.PackSpec) error

type GetPackSpecsFunc func() ([]*fleet.PackSpec, error)

type GetPackSpecFunc func(name string) (*fleet.PackSpec, error)

type NewPackFunc func(pack *fleet.Pack, opts ...fleet.OptionalArg) (*fleet.Pack, error)

type SavePackFunc func(pack *fleet.Pack) error

type DeletePackFunc func(name string) error

type PackFunc func(pid uint) (*fleet.Pack, error)

type ListPacksFunc func(opt fleet.PackListOptions) ([]*fleet.Pack, error)

type PackByNameFunc func(name string, opts ...fleet.OptionalArg) (*fleet.Pack, bool, error)

type ListPacksForHostFunc func(hid uint) (packs []*fleet.Pack, err error)

type EnsureGlobalPackFunc func() (*fleet.Pack, error)

type EnsureTeamPackFunc func(teamID uint) (*fleet.Pack, error)

type PackStore struct {
	ApplyPackSpecsFunc        ApplyPackSpecsFunc
	ApplyPackSpecsFuncInvoked bool

	GetPackSpecsFunc        GetPackSpecsFunc
	GetPackSpecsFuncInvoked bool

	GetPackSpecFunc        GetPackSpecFunc
	GetPackSpecFuncInvoked bool

	NewPackFunc        NewPackFunc
	NewPackFuncInvoked bool

	SavePackFunc        SavePackFunc
	SavePackFuncInvoked bool

	DeletePackFunc        DeletePackFunc
	DeletePackFuncInvoked bool

	PackFunc        PackFunc
	PackFuncInvoked bool

	ListPacksFunc        ListPacksFunc
	ListPacksFuncInvoked bool

	PackByNameFunc        PackByNameFunc
	PackByNameFuncInvoked bool

	ListPacksForHostFunc        ListPacksForHostFunc
	ListPacksForHostFuncInvoked bool

	EnsureGlobalPackFunc        EnsureGlobalPackFunc
	EnsureGlobalPackFuncInvoked bool

	EnsureTeamPackFunc        EnsureTeamPackFunc
	EnsureTeamPackFuncInvoked bool
}

func (s *PackStore) ApplyPackSpecs(specs []*fleet.PackSpec) error {
	s.ApplyPackSpecsFuncInvoked = true
	return s.ApplyPackSpecsFunc(specs)
}

func (s *PackStore) GetPackSpecs() ([]*fleet.PackSpec, error) {
	s.GetPackSpecsFuncInvoked = true
	return s.GetPackSpecsFunc()
}

func (s *PackStore) GetPackSpec(name string) (*fleet.PackSpec, error) {
	s.GetPackSpecFuncInvoked = true
	return s.GetPackSpecFunc(name)
}

func (s *PackStore) NewPack(pack *fleet.Pack, opts ...fleet.OptionalArg) (*fleet.Pack, error) {
	s.NewPackFuncInvoked = true
	return s.NewPackFunc(pack, opts...)
}

func (s *PackStore) SavePack(pack *fleet.Pack) error {
	s.SavePackFuncInvoked = true
	return s.SavePackFunc(pack)
}

func (s *PackStore) DeletePack(name string) error {
	s.DeletePackFuncInvoked = true
	return s.DeletePackFunc(name)
}

func (s *PackStore) Pack(pid uint) (*fleet.Pack, error) {
	s.PackFuncInvoked = true
	return s.PackFunc(pid)
}

func (s *PackStore) ListPacks(opt fleet.PackListOptions) ([]*fleet.Pack, error) {
	s.ListPacksFuncInvoked = true
	return s.ListPacksFunc(opt)
}

func (s *PackStore) PackByName(name string, opts ...fleet.OptionalArg) (*fleet.Pack, bool, error) {
	s.PackByNameFuncInvoked = true
	return s.PackByNameFunc(name, opts...)
}

func (s *PackStore) ListPacksForHost(hid uint) (packs []*fleet.Pack, err error) {
	s.ListPacksForHostFuncInvoked = true
	return s.ListPacksForHostFunc(hid)
}

func (s *PackStore) EnsureGlobalPack() (*fleet.Pack, error) {
	s.EnsureGlobalPackFuncInvoked = true
	return s.EnsureGlobalPackFunc()
}

func (s *PackStore) EnsureTeamPack(teamID uint) (*fleet.Pack, error) {
	s.EnsureTeamPackFuncInvoked = true
	return s.EnsureTeamPackFunc(teamID)
}
