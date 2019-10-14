package cache

import (
	"os"
	"testing"
	"time"
)

var testKey = "testKey"

func TestMain(m *testing.M) {
	if _, err := os.Stat(fileCacheDir); os.IsNotExist(err) {
		if err := os.Mkdir(fileCacheDir, 0644); err != nil {
			panic(err)
		}
	}

	ret := m.Run()

	if err := os.RemoveAll(fileCacheDir); err != nil {
		panic(err)
	}

	os.Exit(ret)
}

func TestMemoryCache_Get(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name string
		fake *MemoryCache
		want bool
	}{
		{
			name: "exist cache",
			fake: &MemoryCache{data: map[string][]byte{testKey: []byte("hoge")}, expires: now.Add(60 * time.Second).Unix()},
			want: true,
		},
		{
			name: "cache expired",
			fake: &MemoryCache{data: map[string][]byte{testKey: []byte("hoge")}, expires: now.Add(-60 * time.Second).Unix()},
			want: false,
		},
		{
			name: "cache not exist",
			fake: &MemoryCache{data: nil, expires: now.Add(-60 * time.Second).Unix()},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fake.Get(testKey)
			if (got != nil) != tt.want {
				t.Fatalf("failed to get cache. want is %v, but got != nil is %v", tt.want, got != nil)
			}
		})
	}
}

func TestMemoryCache_Set(t *testing.T) {
	tests := []struct {
		name string
		arg  []byte
		want bool
	}{
		{name: "success to set cache", arg: []byte("hoge"), want: false},
		{name: "failed to set for existing no cache", arg: nil, want: true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			c := NewMemoryCache(time.Now().Add(DefaultMemoryCacheExpires * time.Second))
			err := c.Set(testKey, tt.arg)
			if (err != nil) != tt.want {
				t.Fatalf("failed to set cache. err is %v but wantErr is %v", err, tt.want)
			}
		})
	}
}

func TestFileCache_Set(t *testing.T) {
	tests := []struct {
		name    string
		arg     []byte
		wantErr bool
	}{
		{name: "success to set cache", arg: []byte("hoge"), wantErr: false},
		{name: "success to overwrite cache", arg: []byte("fuga"), wantErr: false},
		{name: "failed to set cache for empty data", arg: nil, wantErr: true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			fc := &FileCache{}
			err := fc.Set(testKey, tt.arg)
			if (err != nil) != tt.wantErr {
				t.Fatalf("failed to set cache. err is %v but wantErr is %v", err, tt.wantErr)
			}

			got := fc.Get(testKey)
			if tt.arg != nil && string(got) != string(tt.arg) {
				t.Fatalf("failed to set or overwrite cache. got is %v but set is %v", string(got), string(tt.arg))
			}
		})
	}
}

func TestFileCache_Get(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want []byte
	}{
		{name: "success to get cache", key: testKey, want: []byte("hoge")},
		{name: "failed to get cache for invalid key", key: "hoge", want: nil},
	}

	fc := &FileCache{}
	if err := fc.Set(testKey, []byte("hoge")); err != nil {
		t.Fatalf("err: %v", err)
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := fc.Get(tt.key)
			if string(got) != string(tt.want) {
				t.Fatalf("want is %v but got is %v", tt.want, got)
			}
		})
	}
}
