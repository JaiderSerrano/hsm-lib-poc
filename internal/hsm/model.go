package hsm

type ARQCParams struct {
	PAN         string `json:"pan"`
	PSN         string `json:"seqNumber"`
	ATC         string `json:"atc"`
	ARQCMessage string `json:"arqcMessage"`
	ARQC        string `json:"arqc"`
}
