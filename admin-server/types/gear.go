package types

type GearCategoryInType struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GearInType struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

type GearLinkInType struct {
	VendorName  string `json:"vendor_name"`
	Link        string `json:"link"`
	Name        string `json:"name"`
	IsAffiliate bool   `json:"is_affiliate"`
}

type GearCategoryType struct {
	ID          uint32     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Gear        []GearType `json:"gear"`
}

type GearType struct {
	ID          uint32         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	ImageURL    string         `json:"image_url"`
	Links       []GearLinkType `json:"links"`
}

type GearLinkType struct {
	ID          uint32 `json:"id"`
	VendorName  string `json:"vendor_name"`
	Link        string `json:"link"`
	Name        string `json:"name"`
	IsAffiliate bool   `json:"is_affiliate"`
}
