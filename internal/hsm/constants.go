package hsm

import "time"

const (
	appName = "my_bari-hsm-lib-app"
	timeout = 1 * time.Second

	imk      = "828FCB9D68C288F92C78170CF3C22328"
	pek      = "LBKBLKPEKBARIVISA"
	pekID    = "29A4EB53A10B6815B0A57A21DE4D9A10"
	pvkLeft  = "2315208C9110AD40"
	pvkRight = "15EA4CA20131C2FD"
	pvk      = pvkLeft + pvkRight
	cvka     = "E055C1E301E79EEF"
	cvkb     = "A8855D11EE9DD3CE"
)
