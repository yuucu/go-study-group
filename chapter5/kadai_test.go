package chapter5

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   cutEntity
		isError bool
	}{
		{
			name: "成功",
			input: cutEntity{
				args:      []string{"test1"},
				delimiter: ",",
				fields:    1,
			},
			isError: false,
		},
		{
			name: "要素なし",
			input: cutEntity{
				args:      nil,
				delimiter: ",",
				fields:    1,
			},
			isError: true,
		},
		{
			name: "該当データなし",
			input: cutEntity{
				args:      []string{"test3"},
				delimiter: ",",
				fields:    -1,
			},
			isError: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.isError {
				assert.Error(t, Validate(test.input))
			} else {
				assert.NoError(t, Validate(test.input))
			}
		})
	}
}

func TestCut(t *testing.T) {

	type input struct {
		cutEntity
		str io.Reader
	}

	tests := []struct {
		name    string
		input   input
		expect  string
		isError bool
	}{
		{
			name: "成功",
			input: input{
				cutEntity: cutEntity{
					args:      []string{"test1"},
					delimiter: "/",
					fields:    1,
				},
				str: bytes.NewBufferString("aaaaa/bbbbb"),
			},
			expect:  "aaaaa",
			isError: false,
		},
		{
			name: "要素なし",
			input: input{
				cutEntity: cutEntity{
					args:      []string{"test2"},
					delimiter: ",",
					fields:    3,
				},
				str: bytes.NewBufferString("aaaaa/bbbbb"),
			},
			expect:  "",
			isError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var out bytes.Buffer
			if test.isError {
				assert.Error(t, Cut(test.input.cutEntity, test.input.str, &out))
			} else {
				assert.NoError(t, Cut(test.input.cutEntity, test.input.str, &out))
				assert.Equal(t, test.expect, out.String())
			}
		})
	}
}
