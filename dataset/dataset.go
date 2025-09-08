package dataset

import _ "embed"

var Movies = DataSet{
	Name: "movies",
	Catalog: Catalog{
		Cursor:   []string{"film"},
		Required: []string{"year", "film", "publisher", "director", "distributor"},
		Properties: map[string]any{
			"year":            map[string]any{"type": "number"},
			"film":            map[string]any{"type": "string"},
			"publisher":       map[string]any{"type": "string"},
			"director":        map[string]any{"type": "string"},
			"distributor":     map[string]any{"type": "string"},
			"worldwide_gross": map[string]any{"type": "string"},
		},
	},
	Records: []map[string]any{
		{"year": 2020, "film": "Sonic the Hedgehog", "publisher": "Sega Sammy Group", "director": "Jeff Fowler", "distributor": "Paramount Pictures", "worldwide_gross": "$320,954,026"},
		{"year": 2022, "film": "Sonic the Hedgehog 2", "publisher": "Sega Sammy Group", "director": "Jeff Fowler", "distributor": "Paramount Pictures", "worldwide_gross": "$405,421,518"},
		{"year": 2024, "film": "Sonic the Hedgehog 3", "publisher": "Sega Sammy Group", "director": "Jeff Fowler", "distributor": "Paramount Pictures", "worldwide_gross": "$491,603,986"},
	},
}

var Games = DataSet{
	Name: "games",
	Catalog: Catalog{
		Cursor:   []string{"player"},
		Required: []string{"player", "lrt", "date", "platform", "vault_save", "version_number"},
		Properties: map[string]any{
			"player":         map[string]any{"type": "string"},
			"lrt":            map[string]any{"type": "string"},
			"rta":            map[string]any{"type": "string"},
			"date":           map[string]any{"type": "string"},
			"platform":       map[string]any{"type": "string"},
			"vault_save":     map[string]any{"type": "boolean"},
			"version_number": map[string]any{"type": "number"},
		},
	},
	Records: []map[string]any{
		{"player": "wyrnicus", "lrt": "5m 23s 130ms", "rta": "6m 36s 825ms", "date": "2025-07-05", "platform": "pc", "vault_save": false, "version_number": 3420},
		{"player": "lacunae", "lrt": "5m 27s 465ms", "date": "2025-06-30", "platform": "pc", "vault_save": false, "version_number": 3420},
		{"player": "SiDious", "lrt": "5m 39s 960ms", "rta": "6m 09s 905ms", "date": "2025-02-28", "platform": "pc", "vault_save": false, "version_number": 4104},
		{"player": "CantEven", "lrt": "5m 49s 875ms", "rta": "6m 33s 752ms", "date": "2024-02-28", "platform": "pc", "vault_save": false, "version_number": 3420},
		{"player": "Sarahspeedrun", "lrt": "5m 50s 340ms", "rta": "6m 36s 270ms", "date": "2023-07-20", "platform": "pc", "vault_save": false, "version_number": 3420},
		{"player": "Ethan29", "lrt": "5m 56s 580ms", "date": "2022-07-25", "platform": "pc", "vault_save": false, "version_number": 3420},
		{"player": "helyon", "lrt": "5m 56s 775ms", "date": "2024-02-01", "platform": "pc", "vault_save": false, "version_number": 3420},
		{"player": "Shizzai", "lrt": "5m 57s 360ms", "date": "2022-04-05", "platform": "pc", "vault_save": false, "version_number": 3420},
		{"player": "frood", "lrt": "6m 06s 900ms", "rta": "6m 48s 119ms", "date": "2025-06-06", "platform": "pc", "vault_save": false, "version_number": 3420},
		{"player": "KnightedNave", "lrt": "6m 09s 195ms", "rta": "6m 57s 480ms", "date": "2024-12-31", "platform": "pc", "vault_save": false, "version_number": 4104},
	},
}

var Custom = DataSet{
	Name: "custom",
	Catalog: Catalog{
		Cursor:     []string{},
		Required:   []string{},
		Properties: map[string]any{},
	},
	Records: nil,
}

type DataSet struct {
	Name    string
	Catalog Catalog
	Records Records
}

type Catalog struct {
	Cursor     []string       `json:"cursor"`
	Required   []string       `json:"required"`
	Properties map[string]any `json:"properties"`
}

type Records []map[string]any
