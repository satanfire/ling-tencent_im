package tencentim

import "testing"

func TestSend(t *testing.T) {
	SetAppID("14002")
	sig := "eJxNjF1Pgzkk6mjM1kC0lT1jLqwsfaIpvG-y4hW-T2nOc832j1kt3y3a7tG8vsuZPoHgG6mbASsrGqVFKPkItaNRfBu04Jxi1ztfi3N*LAJjUyTAEIkDtKLlKeOqUl46Wd7rDneQTgmn5KbVTbjIIA9jBxAf6kVbWcEh8o9oPgemnUfsRJuHmM0yflPs-Iscqd6j3ZDNmbJYMu6eFLFEG4CGnuLKNzsj8Wyy3Mh7h6WPQ8pgVfz01rXsO0jixUq8bvozp3fAjWpzQCDbOt-egz9PMLE1FYvA__"
	NewSendMsg(sig, "USERxxxx", "xxx")
}

func BenchmarkTestSend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetAppID("1400202")
		sig := "eJxNjtFOgzAYRt*FW43*Lddtpsy47tLyBa9PSfny-ftrBfpHS*K7tja3PZKODMHnNsRy1K0VlZS6AFmaZDMXPxALpIrJcuc29zV5b-GlE0*qoEhAoABTwi*SPGlpBY5r*w4iSilGOCanoQ2smsHgQFRhF2AP2nlhxgTBgQjRq83jKwHHAeZH3mknhum4*jlEMaHqrb6HAaKr8Bt-EkinvYLvFuiPd1tl54MPBOVfcq6sNH9ukeM8eKdoCNNzEat4rf7af15uslC*-w696ePzs8v25JX7A__"
		NewSendMsg(sig, ":3294", "xxx")
	}
}
