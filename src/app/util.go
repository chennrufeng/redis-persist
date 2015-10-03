package main

// 写DB后的，索引前缀
const INDEX_KEY_PREFIX string = "|"

// DB中的索引前缀长度
const INDEX_KEY_LEN int = len("|")

// DB中索引的开始字节
var INDEX_KEY_START = []byte("|")

// DB中索引的结束字节
var INDEX_KEY_END = []byte{'|', 0xff}

// DB中KEY开始字节
var KEY_START = []byte("uid:")

// DB中KEY结束字节
var KEY_END = []byte{'u', 'i', 'd', ':', 0xff}

func indexKey(key string) string {
	return INDEX_KEY_PREFIX + key
}
