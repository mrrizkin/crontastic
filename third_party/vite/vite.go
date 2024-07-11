package vite

import (
	goVite "github.com/mrrizkin/go-vite-parser"
)

func New() goVite.ViteManifestInfo {
	return goVite.Parse(goVite.Config{
		OutDir:       "/build/",
		ManifestPath: "public/build/manifest.json",
		HotFilePath:  "public/hot",
	})
}
