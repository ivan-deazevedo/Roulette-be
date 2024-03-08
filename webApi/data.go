package webapi

type Restaurants struct {
	ID           int    `json:"id"`
	Naam         string `json:"naam"`
	Omschrijving string `json:"omschrijving"`
}

type Opties []Restaurants

var Data = Opties{
	{ID: 1, Naam: "Subway", Omschrijving: "Broodjes zaak"},
	{ID: 2, Naam: "Khun Kaew", Omschrijving: "Thais"},
	{ID: 3, Naam: "Dominos", Omschrijving: "Pizza"},
	{ID: 4, Naam: "Minh Anh", Omschrijving: "Vietnamees"},
	{ID: 5, Naam: "Pat's Tosti Bar", Omschrijving: "Broodjes zaak"},
	{ID: 6, Naam: "Cho Pain", Omschrijving: "Broodjes zaak"},
}
