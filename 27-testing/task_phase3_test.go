package main

import (
	"testing"

	"github.com/kulti/titlecase"
	"github.com/stretchr/testify/assert"
)

func TestTdt(t *testing.T) {
	testCases := []struct {
		desc             string
		str, minor, want string
	}{
		{
			desc:  "Empty",
			str:   "",
			minor: "",
			want:  "",
		},
		{
			desc:  "WithoutMinor",
			str:   "the quick fox in the bag",
			minor: "",
			want:  "The Quick Fox In The Bag",
		},
		{
			desc:  "WithMinorInFirst",
			str:   "the quick fox in the bag",
			minor: "the",
			want:  "The Quick Fox In the Bag",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := titlecase.TitleCase(tC.str, tC.minor)
			assert.Equal(t, tC.want, got)
		})
	}
}
