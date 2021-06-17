package server_chan

import "testing"

func TestNotifier_Notify(t *testing.T) {
	n := NewNotifier("your key")
	_ = n.Notify("haha", "hfue")
}
