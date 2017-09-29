package sigqueue

import "testing"

func Test_Sigqueue_Create(t *testing.T) {
	sq := CreateSigqueue()
	if sq == nil {
		t.Error("Got nil")
	}
}
