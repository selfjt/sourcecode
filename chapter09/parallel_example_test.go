package main

import "testing"

var pairs = []struct {
	k string
	v string
}{
	{"polaris", " 徐新华 "},
	{"studygolang", "Go 语言中文网 "},
	{"stdlib", "Go 语言标准库 "},
	{"polaris1", " 徐新华 1"},
	{"studygolang1", "Go 语言中文网 1"},
	{"stdlib1", "Go 语言标准库 1"},
	{"polaris2", " 徐新华 2"},
	{"studygolang2", "Go 语言中文网 2"},
	{"stdlib2", "Go 语言标准库 2"},
	{"polaris3", " 徐新华 3"},
	{"studygolang3", "Go 语言中文网 3"},
	{"stdlib3", "Go 语言标准库 3"},
	{"polaris4", " 徐新华 4"},
	{"studygolang4", "Go 语言中文网 4"},
	{"stdlib4", "Go 语言标准库 4"},
}
// TODO 我们测试时 需要先写入 然后在读取 不然起不到验证效果的作用
func TestWriteToMap(t *testing.T) {
	// 增加并行写入
	t.Parallel()
	for _, tt := range pairs {
		WriteToMap(tt.k, tt.v)
	}
}

// 读取map
func TestReaderToMap(t *testing.T) {
	t.Parallel()
	for _, tt := range pairs {
		v := ReaderToMap(tt.k)
		if v != tt.v {
			t.Errorf("the value of key(%s) is %s, expected: %s", tt.k, v, tt.v)
		}
	}
}
