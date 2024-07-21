package main

import (
	"os"
	"log"
)

func loadFont(fontPath string) []byte {
	// 加载字体文件
	fontData, err := os.ReadFile(fontPath)
	if err != nil {
		log.Fatalf("无法加载字体文件: %v", err)
	}
	return fontData
}