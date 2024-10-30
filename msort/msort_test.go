package msort

import "testing"

func TestMsort(t *testing.T) {
	var assorted = [...]string{
		"delta",
		"echo",
		"lima",
		"alpha",
		"bravo",
		"romeo",
		"uniform",
		"hotel",
		"zulu",
		"india",
		"yankee",
		"oscar",
		"charlie",
		"golf",
		"victor",
		"papa",
		"kilo",
		"mike",
		"sierra",
		"tango",
		"juliet",
		"x-ray",
		"quebec",
		"foxtrot",
		"whiskey",
		"november",
	}
	var want = [...]string{
		"alpha",
		"bravo",
		"charlie",
		"delta",
		"echo",
		"foxtrot",
		"golf",
		"hotel",
		"india",
		"juliet",
		"kilo",
		"lima",
		"mike",
		"november",
		"oscar",
		"papa",
		"quebec",
		"romeo",
		"sierra",
		"tango",
		"uniform",
		"victor",
		"whiskey",
		"x-ray",
		"yankee",
		"zulu",
	}
	var sorted [26]string
	copy(sorted[0:], assorted[0:])
	Msort(sorted[0:])
	if sorted != want {
		t.Fatalf("Failed to sort strings")
	}
}
