package pick_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	log "github.com/sirupsen/logrus"

	"github.com/asphaltbuffet/go-picker/pkg/pick"
)

func TestPickOne(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	pkr := pick.NewPicker()

	type args struct {
		indexer func(int) int
		input   []any
	}

	var (
		colors = []any{"red", "blue", "green", "yellow"}
		nums   = []any{1, 2, 3, 4, 5, 6}
		mixed  = []any{1, "red", 3, "blue", 5, "green"}
	)

	tests := []struct {
		name         string
		args         args
		want         any
		errAssertion assert.ErrorAssertionFunc
		wantErr      error
	}{
		{"strings - first element", args{indexer: firstIndex, input: colors}, "red", assert.NoError, nil},
		{"strings - last element", args{indexer: lastIndex, input: colors}, "yellow", assert.NoError, nil},
		{"ints - last element", args{indexer: lastIndex, input: nums}, 6, assert.NoError, nil},
		{"mixed - first", args{indexer: firstIndex, input: mixed}, 1, assert.NoError, nil},
		{"mixed - last", args{indexer: lastIndex, input: mixed}, "green", assert.NoError, nil},
		{"empty", args{indexer: firstIndex, input: []any{}}, nil, assert.Error, pick.ErrEmptyCollection},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pkr.Rand = tt.args.indexer

			got, err := pick.GetOne(pkr, tt.args.input)
			tt.errAssertion(t, err)

			if err != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func firstIndex(i int) int {
	return 0
}

func lastIndex(i int) int {
	return i - 1
}
