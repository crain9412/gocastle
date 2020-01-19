package parallel

import (
	"testing"
	"time"
)

func Test_SafeMap(t *testing.T) {
	want := []string{"the", "quick", "fox", "jumps", "over", "the", "lazy", "dog"}
	want2 := []string{"the", "dog", "runs"}
	documentName := "My Document.docx"
	safeMap := InitSafeMap()
	go safeMap.Put(documentName, want)
	go safeMap.Put(documentName, want2)
	time.Sleep(100 * time.Millisecond)
	got, ok := safeMap.Get(documentName)

	if !ok {
		t.Errorf("Map didn't contain any values")
	}

	for i := 0; i < len(got); i++ {
		if i < len(want) && got[i] != want[i] && i < len(want2) && got[i] != want2[i] {
			t.Errorf("String %d was %s, want %s or %s", i, got, want, want2)
		}
	}
}
