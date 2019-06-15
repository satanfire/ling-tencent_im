package tencentim

import "testing"

func TestSend(t *testing.T) {
	SetAppID("142742")
	sig := ""
	sendMsg := NewSendMsg(sig, "USER:", "xxx")
	send, err := Send(sendMsg)
	if err != nil {
		t.Fatal("test error", err)
	}
	t.Log(string(send))
}

func BenchmarkTestSend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetAppID("1400202")
		sig := ""
		sendMsg := NewSendMsg(sig, "", "xxx")
		send, err := Send(sendMsg)
		if err != nil {
			b.Fatal("test error", err)
		}
		b.Log(string(send))
	}
}
