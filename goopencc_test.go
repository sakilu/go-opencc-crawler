package goopencc

import (
	"testing"
)

func TestZh2Tw(t *testing.T) {
	s, statusCode := Zh2Tw("鼠标")
	if s != "滑鼠" || statusCode != 200 {
		t.Error(statusCode, ":", s, "!=", "滑鼠")
	}
	s, statusCode = Zh2Tw(" ")
	if s != "" || statusCode != 200 {
		t.Error("Zh2Tw:", "空白測試失敗")
	}
}

func TestTw2Zh(t *testing.T) {
	s, statusCode := Tw2Zh("滑鼠")
	if s != "鼠标" || statusCode != 200 {
		t.Error(statusCode, ":", s, "!=", "鼠标")
	}
	s, statusCode = Tw2Zh(" ")
	if s != "" || statusCode != 200 {
		t.Error("Tw2Zh:", "空白測試失敗")
	}
}
