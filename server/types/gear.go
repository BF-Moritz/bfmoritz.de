package types

type GearCategoryType struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Gear        []GearType `json:"gear"`
}

type GearType struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	ImageURL    string         `json:"image_url"`
	Links       []GearLinkType `json:"links"`
}

type GearLinkType struct {
	VendorName  string `json:"vendor_name"`
	Link        string `json:"link"`
	Name        string `json:"name"`
	IsAffiliate bool   `json:"is_affiliate"`
}
