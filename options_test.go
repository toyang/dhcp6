package dhcp6

import (
	"bytes"
	"reflect"
	"testing"
)

// TestOptionsAdd verifies that Options.Add correctly creates or appends
// key/value Option pairs to an Options map.
func TestOptionsAdd(t *testing.T) {
	var tests = []struct {
		description string
		kv          []Option
		options     Options
	}{
		{
			description: "one key/value pair",
			kv: []Option{
				Option{
					Code: 1,
					Data: []byte("foo"),
				},
			},
			options: Options{
				1: [][]byte{[]byte("foo")},
			},
		},
		{
			description: "two key/value pairs",
			kv: []Option{
				Option{
					Code: 1,
					Data: []byte("foo"),
				},
				Option{
					Code: 2,
					Data: []byte("bar"),
				},
			},
			options: Options{
				1: [][]byte{[]byte("foo")},
				2: [][]byte{[]byte("bar")},
			},
		},
		{
			description: "three key/value pairs, two with same key",
			kv: []Option{
				Option{
					Code: 1,
					Data: []byte("foo"),
				},
				Option{
					Code: 1,
					Data: []byte("baz"),
				},
				Option{
					Code: 2,
					Data: []byte("bar"),
				},
			},
			options: Options{
				1: [][]byte{[]byte("foo"), []byte("baz")},
				2: [][]byte{[]byte("bar")},
			},
		},
	}

	for i, tt := range tests {
		o := make(Options)
		for _, p := range tt.kv {
			o.Add(p.Code, p.Data)
		}

		if want, got := tt.options, o; !reflect.DeepEqual(want, got) {
			t.Fatalf("[%02d] test %q, unexpected Options map:\n- want: %v\n-  got: %v",
				i, tt.description, want, got)
		}
	}
}

// TestOptionsGet verifies that Options.Get correctly selects the first value
// for a given key, if the value is not empty in an Options map.
func TestOptionsGet(t *testing.T) {
	var tests = []struct {
		description string
		options     Options
		key         OptionCode
		value       []byte
		ok          bool
	}{
		{
			description: "nil Options map",
		},
		{
			description: "empty Options map",
			options:     Options{},
		},
		{
			description: "value not present in Options map",
			options: Options{
				2: [][]byte{[]byte("foo")},
			},
			key: 1,
		},
		{
			description: "value present in Options map, but zero length value for key",
			options: Options{
				1: [][]byte{},
			},
			key: 1,
		},
		{
			description: "value present in Options map",
			options: Options{
				1: [][]byte{[]byte("foo")},
			},
			key:   1,
			value: []byte("foo"),
			ok:    true,
		},
		{
			description: "value present in Options map, with multiple values",
			options: Options{
				1: [][]byte{[]byte("foo"), []byte("bar")},
			},
			key:   1,
			value: []byte("foo"),
			ok:    true,
		},
	}

	for i, tt := range tests {
		value, ok := tt.options.Get(tt.key)

		if want, got := tt.value, value; !bytes.Equal(want, got) {
			t.Fatalf("[%02d] test %q, unexpected value for Options.Get(%v):\n- want: %v\n-  got: %v",
				i, tt.description, tt.key, want, got)
		}

		if want, got := tt.ok, ok; want != got {
			t.Fatalf("[%02d] test %q, unexpected ok for Options.Get(%v): %v != %v",
				i, tt.description, tt.key, want, got)
		}
	}
}

// TestOptionsClientID verifies that Options.ClientID properly parses and returns
// a DUID value, if one is available with OptionClientID.
func TestOptionsClientID(t *testing.T) {
	var tests = []struct {
		description string
		options     Options
		duid        DUID
		ok          bool
	}{
		{
			description: "OptionClientID not present in Options map",
		},
		{
			description: "OptionClientID present in Options map",
			options: Options{
				OptionClientID: [][]byte{[]byte{0, 1}},
			},
			duid: DUIDLLT([]byte{0, 1}),
			ok:   true,
		},
	}

	for i, tt := range tests {
		// DUID parsing is tested elsewhere, so errors should automatically fail
		// test here
		duid, ok, err := tt.options.ClientID()
		if err != nil {
			t.Fatal(err)
		}

		if want, got := tt.duid, duid; !reflect.DeepEqual(want, got) {
			t.Fatalf("[%02d] test %q, unexpected value for Options.ClientID():\n- want: %v\n-  got: %v",
				i, tt.description, want, got)
		}

		if want, got := tt.ok, ok; want != got {
			t.Fatalf("[%02d] test %q, unexpected ok for Options.ClientID(): %v != %v",
				i, tt.description, want, got)
		}
	}
}

// TestOptionsServerID verifies that Options.ServerID properly parses and returns
// a DUID value, if one is available with OptionServerID.
func TestOptionsServerID(t *testing.T) {
	var tests = []struct {
		description string
		options     Options
		duid        DUID
		ok          bool
	}{
		{
			description: "OptionServerID not present in Options map",
		},
		{
			description: "OptionServerID present in Options map",
			options: Options{
				OptionServerID: [][]byte{[]byte{0, 1}},
			},
			duid: DUIDLLT([]byte{0, 1}),
			ok:   true,
		},
	}

	for i, tt := range tests {
		// DUID parsing is tested elsewhere, so errors should automatically fail
		// test here
		duid, ok, err := tt.options.ServerID()
		if err != nil {
			t.Fatal(err)
		}

		if want, got := tt.duid, duid; !reflect.DeepEqual(want, got) {
			t.Fatalf("[%02d] test %q, unexpected value for Options.ServerID():\n- want: %v\n-  got: %v",
				i, tt.description, want, got)
		}

		if want, got := tt.ok, ok; want != got {
			t.Fatalf("[%02d] test %q, unexpected ok for Options.ServerID(): %v != %v",
				i, tt.description, want, got)
		}
	}
}
