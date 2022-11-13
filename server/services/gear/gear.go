package gear

import (
	"server/daos/gear"
	"server/types"
	"server/vars"
	"sort"
	"strings"
)

type Service struct {
	gearDAO gear.DAOInterface
}

type ServiceInterface interface {
	GetGear() (gear []types.GearCategoryType, err error)
}

func NewService() ServiceInterface {
	return &Service{
		gearDAO: gear.NewDAO(),
	}
}

func (s *Service) GetGear() (gear []types.GearCategoryType, err error) {

	// get categories
	gear, err = s.gearDAO.GetAllGear()
	if err != nil {
		vars.Logger.LogError("service:gear.GetGear()", "failed to get all gear (%s)", err)
		return nil, err
	}

	// sort categories by name
	sort.Slice(gear, func(i, j int) bool {
		return strings.Compare(gear[i].Name, gear[j].Name) < 0
	})

	// sort gear by name
	for g := range gear {
		sort.Slice(gear[g].Gear, func(i, j int) bool {
			return strings.Compare(gear[g].Gear[i].Name, gear[g].Gear[j].Name) < 0
		})

		// sort links by venor name
		for h := range gear[g].Gear {
			sort.Slice(gear[g].Gear[h].Links, func(i, j int) bool {
				return strings.Compare(gear[g].Gear[h].Links[i].VendorName, gear[g].Gear[h].Links[j].VendorName) < 0
			})
		}
	}

	return
}
