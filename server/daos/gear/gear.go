package gear

import (
	"server/types"
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
