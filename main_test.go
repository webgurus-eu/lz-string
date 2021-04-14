package lzstring

import (
	"testing"
)

var testingValues = []string{"alma", "körte", "{\"id\":\"1234534625254\",\"at\":2,\"tmax\":120,\"imp\":[{\"id\":\"1\",\"banner\":{\"w\":1583,\"h\":1095,\"pos\":7,\"battr\":[13]}}, {\"id\":\"2\",\"banner\":{\"w\":784,\"h\":100,\"pos\":4,\"battr\":[13]}}],\"badv\":[\"company1.com\",\"company2.com\"],\"site\":{\"id\":\"234563\",\"name\":\"Site ABCD\",\"domain\":\"www.auchandrive.fr\",\"pagecat\":[\"3700610\"],\"sectioncat\":[\"3700624\"],\"privacypolicy\":1,\"page\":\"http://siteabcd.com/page.htm\",\"ref\":\"http://referringsite.com/referringpage.htm\",\"publisher\":{\"id\":\"pub12345\",\"name\":\"Publisher A\"},\"content\":{\"keywords\":[\"keyword a\",\"keyword b\",\"keyword c\"]}},\"device\":{\"ip\":\"64.124.253.1\",\"ua\":\"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.2.16) Gecko/20110319 Firefox/3.6.16\",\"os\":\"OS X\",\"flashver\":\"10.1\",\"js\":1},\"user\":{\"id\":\"45asdf987656789adfad4678rew656789\",\"buyeruid\":\"5df678asd8987656asdf78987654\",\"data\":[{\"name\":\"AuchanDrive\", \"segment\":[{\"name\":\"basket\", \"value\":\"[]\"}]}]}}"}

func TestDecompress(t *testing.T) {
	for _, k := range testingValues {
		encoded := CompressToBase64(k)

		result, err := DecompressFromBase64(encoded)
		if err != nil {
			t.Error("Unexpected error", err)
		}
		if result != k {
			t.Error("Result should be :\n", k, "\n instead of :\n", result)
		}
	}
}

func BenchmarkDecompress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, k := range testingValues {
			DecompressFromBase64(k)
		}
	}
}
