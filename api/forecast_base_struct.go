package api

type ForecastBaseResponse struct {
	Localita Localita `json:"localita"`
}
type Vento struct {
	Direzione string `json:"direzione"`
	Intensita string `json:"intensita"`
}
type AllertePrevisioni struct {
	P float64 `json:"p"`
	N float64 `json:"n"`
	T float64 `json:"t"`
	V float64 `json:"v"`
	G float64 `json:"g"`
}
type TempoMedio struct {
	IDSimbolo         string            `json:"id_simbolo"`
	TMin              int               `json:"t_min"`
	TMax              int               `json:"t_max"`
	Vento             Vento             `json:"vento"`
	Mare              interface{}       `json:"mare"`
	Hr                float64           `json:"hr"`
	Pr                float64           `json:"pr"`
	Accumulo          interface{}       `json:"accumulo"`
	Precipitazioni    float64           `json:"precipitazioni"`
	PrecInt           string            `json:"prec_int"`
	ProbabilitaPrec   int               `json:"probabilita_prec"`
	NP                string            `json:"n_p"`
	PrecUnita         string            `json:"prec_unita"`
	AllertePrevisioni AllertePrevisioni `json:"allerte_previsioni"`
	Raffica           float64           `json:"raffica"`
	IDSimboloMaso     string            `json:"id_simbolo_maso"`
	Speciale          float64           `json:"speciale"`
}
type Temperatura struct {
	Gradi float64 `json:"gradi"`
	Quota string  `json:"quota"`
}
type PrevisioneOraria struct {
	Ora               float64           `json:"ora"`
	Temperatura       Temperatura       `json:"temperatura"`
	Tpercepita        float64           `json:"tpercepita"`
	Zt                string            `json:"zt"`
	Qn                string            `json:"qn"`
	Vento             Vento             `json:"vento"`
	Mare              interface{}       `json:"mare"`
	Raffica           float64           `json:"raffica"`
	Hr                string            `json:"hr"`
	Pr                float64           `json:"pr"`
	Uv                float64           `json:"uv"`
	Accumulo          interface{}       `json:"accumulo"`
	Notte             float64           `json:"notte"`
	Precipitazioni    float64           `json:"precipitazioni"`
	PrecInt           string            `json:"prec_int"`
	ProbabilitaPrec   float64           `json:"probabilita_prec"`
	NP                string            `json:"n_p"`
	PrecUnita         string            `json:"prec_unita"`
	IDSimbolo         string            `json:"id_simbolo"`
	AllertePrevisioni AllertePrevisioni `json:"allerte_previsioni"`
	MostraFascia      interface{}       `json:"mostra_fascia"`
	Windchill         float64           `json:"windchill"`
	DescBreve         string            `json:"desc_breve"`
	IDSimboloMaso     string            `json:"id_simbolo_maso"`
	Speciale          float64           `json:"speciale"`
}
type PrevisioneGiorno struct {
	Data             string             `json:"data"`
	TempoMedio       TempoMedio         `json:"tempo_medio"`
	PrevisioneOraria []PrevisioneOraria `json:"previsione_oraria,omitempty"`
	Attendibilita    string             `json:"attendibilita"`
	FasceSuccessive  string             `json:"fasce_successive"`
}
type Foto struct {
	ID          string `json:"id"`
	Data        string `json:"data"`
	Titolo      string `json:"titolo"`
	IDLocalita  string `json:"id_localita"`
	Localita    string `json:"localita"`
	FilenameHd  string `json:"filename_hd"`
	ThumbnailHd string `json:"thumbnail_hd"`
}
type Localita struct {
	ID               int                `json:"id"`
	Localita         string             `json:"localita"`
	Prov             string             `json:"prov"`
	PrevisioneGiorno []PrevisioneGiorno `json:"previsione_giorno"`
	IDMacrosettore   float64            `json:"id_macrosettore"`
	Altitudine       float64            `json:"altitudine"`
	Lat              float64            `json:"lat"`
	Lon              float64            `json:"lon"`
	Popolazione      float64            `json:"popolazione"`
	Tipo             string             `json:"tipo"`
	CodiceStato      string             `json:"codice_stato"`
	Impianti         interface{}        `json:"impianti"`
	CodiceIstat      string             `json:"codice_istat"`
	Adv              string             `json:"adv"`
	IDSettore        float64            `json:"id_settore"`
	NeveSi           float64            `json:"neve_si"`
	CanonicalURL     string             `json:"canonical_url"`
	Reporter         interface{}        `json:"reporter"`
	Foto             []Foto             `json:"foto"`
	Webcam           interface{}        `json:"webcam"`
	LocalitaVicine   interface{}        `json:"localita_vicine"`
	PosizioniBanner  string             `json:"posizioni_banner"`
	Nazione          string             `json:"nazione"`
	Regione          string             `json:"regione"`
	IDLocalitaOld    string             `json:"id_localita_old"`
	ProvEstesa       string             `json:"prov_estesa"`
}
