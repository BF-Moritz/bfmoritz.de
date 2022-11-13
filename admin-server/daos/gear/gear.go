package gear

import (
	"admin-server/types"
)

type DAO struct {
}

type DAOInterface interface {
	GetAllGear() (categories []types.GearCategoryType, err error)
	// GetVideosTags() (tags map[string][]string, err error)
	// GetVideosMeta() (meta map[string]types.VideoMetaType, err error)
}

func NewDAO() DAOInterface {
	return &DAO{}
}

func (d *DAO) GetAllGear() (categories []types.GearCategoryType, err error) {
	// TODO
	categories = []types.GearCategoryType{
		{
			Name:        "test",
			Description: "das ist nur 1 test",
			Gear: []types.GearType{
				{
					Name:        "hi",
					Description: "penis",
					ImageURL:    "img.url",
					Links: []types.GearLinkType{
						{
							VendorName:  "b",
							Link:        "",
							Name:        "dawd",
							IsAffiliate: false,
						},
						{
							VendorName:  "akk",
							Link:        "",
							Name:        "dawd",
							IsAffiliate: false,
						},
						{
							VendorName:  "buh",
							Link:        "",
							Name:        "dawd",
							IsAffiliate: false,
						},
					},
				},
				{
					Name:        "hallo",
					Description: "penis",
					ImageURL:    "img.url",
					Links: []types.GearLinkType{
						{
							VendorName:  "t",
							Link:        "",
							Name:        "dawd",
							IsAffiliate: false,
						},
						{
							VendorName:  "typ",
							Link:        "",
							Name:        "dawd",
							IsAffiliate: false,
						},
						{
							VendorName:  "buh",
							Link:        "",
							Name:        "dawd",
							IsAffiliate: false,
						},
					},
				},
			},
		},
		{
			Name:        "est",
			Description: "das ist nur 1 test",
			Gear: []types.GearType{
				{
					Name:        "hi",
					Description: "penis",
					ImageURL:    "img.url",
					Links: []types.GearLinkType{
						{
							VendorName:  "b",
							Link:        "",
							Name:        "dawd",
							IsAffiliate: false,
						},
						{
							VendorName:  "akk",
							Link:        "",
							Name:        "dawd",
							IsAffiliate: false,
						},
						{
							VendorName:  "buh",
							Link:        "",
							Name:        "dawd",
							IsAffiliate: false,
						},
					},
				},
				{
					Name:        "hallo",
					Description: "penis",
					ImageURL:    "img.url",
					Links: []types.GearLinkType{
						{
							VendorName:  "t",
							Link:        "",
							Name:        "dawd",
							IsAffiliate: false,
						},
						{
							VendorName:  "typ",
							Link:        "",
							Name:        "dawd",
							IsAffiliate: false,
						},
						{
							VendorName:  "buh",
							Link:        "",
							Name:        "dawd",
							IsAffiliate: false,
						},
					},
				},
			},
		},
		{
			Name:        "st",
			Description: "das ist nur 1 test",
			Gear: []types.GearType{
				{
					Name:        "hi",
					Description: "penis",
					ImageURL:    "img.url",
					Links: []types.GearLinkType{
						{
							VendorName:  "b",
							Link:        "",
							Name:        "dawd",
							IsAffiliate: false,
						},
						{
							VendorName:  "akk",
							Link:        "",
							Name:        "dawd",
							IsAffiliate: false,
						},
						{
							VendorName:  "buh",
							Link:        "",
							Name:        "dawd",
							IsAffiliate: false,
						},
					},
				},
				{
					Name:        "hallo",
					Description: "penis",
					ImageURL:    "img.url",
					Links: []types.GearLinkType{
						{
							VendorName:  "t",
							Link:        "",
							Name:        "dawd",
							IsAffiliate: false,
						},
						{
							VendorName:  "typ",
							Link:        "",
							Name:        "dawd",
							IsAffiliate: false,
						},
						{
							VendorName:  "buh",
							Link:        "",
							Name:        "dawd",
							IsAffiliate: false,
						},
					},
				},
			},
		},
		{
			Name:        "t",
			Description: "das ist nur 1 test",
			Gear: []types.GearType{
				{
					Name:        "hi",
					Description: "penis",
					ImageURL:    "img.url",
					Links: []types.GearLinkType{
						{
							VendorName:  "b",
							Link:        "",
							Name:        "dawd",
							IsAffiliate: false,
						},
						{
							VendorName:  "akk",
							Link:        "",
							Name:        "dawd",
							IsAffiliate: false,
						},
						{
							VendorName:  "buh",
							Link:        "",
							Name:        "dawd",
							IsAffiliate: false,
						},
					},
				},
				{
					Name:        "hallo",
					Description: "penis",
					ImageURL:    "img.url",
					Links: []types.GearLinkType{
						{
							VendorName:  "t",
							Link:        "",
							Name:        "dawd",
							IsAffiliate: false,
						},
						{
							VendorName:  "typ",
							Link:        "",
							Name:        "dawd",
							IsAffiliate: false,
						},
						{
							VendorName:  "buh",
							Link:        "",
							Name:        "dawd",
							IsAffiliate: false,
						},
					},
				},
			},
		},
	}

	return
}

func (d *DAO) DeleteCategory(id uint32) (err error) {
	// TODO
	return nil
}

func (d *DAO) DeleteGear(id uint32) (err error) {
	// TODO
	return nil
}

func (d *DAO) DeleteLink(id uint32) (err error) {
	// TODO
	return nil
}

func (d *DAO) InsertCategory(cat types.GearCategoryInType) (insertedCat types.GearCategoryType, err error) {
	// TODO
	return
}

func (d *DAO) InsertGear(gear types.GearInType) (insertedGear types.GearType, err error) {
	// TODO
	return
}

func (d *DAO) InsertLink(link types.GearLinkInType) (insertedLink types.GearLinkType, err error) {
	// TODO
	return
}

func (d *DAO) UpdateCategory(id uint32, cat types.GearCategoryInType) (insertedCat types.GearCategoryType, err error) {
	// TODO
	return
}

func (d *DAO) UpdateGear(id uint32, gear types.GearInType) (insertedGear types.GearType, err error) {
	// TODO
	return
}

func (d *DAO) UpdateLink(id uint32, link types.GearLinkInType) (insertedLink types.GearLinkType, err error) {
	// TODO
	return
}
