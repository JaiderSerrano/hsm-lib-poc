package hsm

type ARQCParams struct {
	PAN         string `json:"pan"`
	PSN         string `json:"seqNumber"`
	ATC         string `json:"atc"`
	ARQCMessage string `json:"arqcMessage"`
	ARQC        string `json:"arqc"`
}

type PINGenerationParams struct {
	PAN  string `json:"pan"`
	PVKI string `json:"pvki"`
}

type PVVGenerationParams struct {
	PIN  string `json:"pin"`
	PAN  string `json:"pan"`
	PVKI string `json:"pvki"`
}
