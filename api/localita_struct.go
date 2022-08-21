package api

type LocalitaResponse struct {
	Localita []struct {
		ID            int     `json:"id"`
		Localita      string  `json:"localita"`
		Prov          string  `json:"prov"`
		Lat           float64 `json:"lat"`
		Lon           float64 `json:"lon"`
		Tipo          string  `json:"tipo"`
		CodiceStato   string  `json:"codice_stato"`
		IDSettore     int     `json:"id_settore"`
		CanonicalURL  string  `json:"canonical_url"`
		Nazione       string  `json:"nazione"`
		Regione       string  `json:"regione"`
		IDLocalitaOld string  `json:"id_localita_old"`
		ProvEstesa    string  `json:"prov_estesa"`
		Timezone      string  `json:"timezone"`
	} `json:"localita"`
}
