package main

import "testing"

func TestExtractPath(t *testing.T) {
	str := `[Trash Info]
Path=/mnt/disk/me/temp/fk/orgpress/org/theindex.inc
DeletionDate=2023-02-12T00:18:46`
	want := `/mnt/disk/me/temp/fk/orgpress/org/theindex.inc`
	got := extractPath(str)
	if want != got {
		t.Errorf("wanted %s got %s", want, got)
	}
}
