//////
//  alternate basic generate punk example
//     punk #0 and #1 using basic 24x24 attribute images / building blocks

package main


import (
  "fmt"

	"github.com/learnpixelart/pixelart.go/pixelart"
)


var dir = "../basic"


func main() {
  fmt.Printf( "Hello, Pixel Art v%s!\n", pixelart.Version )

  ///////////
	// read in f(emale) attributes
	female2        := pixelart.ReadImage( dir + "/female2.png" )
  earring        := pixelart.ReadImage( dir + "/f/earring.png" )
  blondebob      := pixelart.ReadImage( dir + "/f/blondebob.png" )
  greeneyeshadow := pixelart.ReadImage( dir + "/f/greeneyeshadow.png" )

	// test drive
	// generate punk #0
	punk := pixelart.NewImage( 24, 24 )
	punk.Paste( female2 )
	punk.Paste( earring )
	punk.Paste( blondebob )
	punk.Paste( greeneyeshadow )

	punk.Save( "./punk0.png" )
	punk.Zoom(20).Save( "./punk0@20x.png" )

  // (re)try with background
	punk = pixelart.NewImage( 24, 24 ).Background( "#60A4F7" )
	punk.Paste( female2 )
	punk.Paste( earring )
	punk.Paste( blondebob )
	punk.Paste( greeneyeshadow )

	punk.Save( "./bluepunk0.png" )
	punk.Zoom(20).Save( "./bluepunk0@20x.png" )


  ///////////
	// read in m(ale) attributes
	male1   := pixelart.ReadImage( dir + "/male1.png" )
  smile   := pixelart.ReadImage( dir + "/m/smile.png" )
  mohawk  := pixelart.ReadImage( dir + "/m/mohawk.png" )

	// generate punk #1
	punk = pixelart.NewImage( 24, 24 )
  punk.Paste( male1 )
	punk.Paste( smile )
	punk.Paste( mohawk )

	punk.Save( "./punk1.png" )
	punk.Zoom(20).Save( "./punk1@20x.png" )

	// (re)try with background
	punk = pixelart.NewImage( 24, 24 ).Background( "#60A4F7" )
  punk.Paste( male1 )
	punk.Paste( smile )
	punk.Paste( mohawk )

	punk.Save( "./bluepunk1.png" )
	punk.Zoom(20).Save( "./bluepunk1@20x.png" )

  fmt.Println( "Bye")
}


