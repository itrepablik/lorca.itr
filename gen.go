//+build generate

package main

import "github.com/itrepablik/lorca.itr"

func main() {
	// You can also run "npm build" or webpack here, or compress assets, or
	// generate manifests, or do other preparations for your assets.

	lorca.Embed("main", "assets.go", "www", "www/js")
	//lorca.Embed("main", "assets.go", "www")
}
