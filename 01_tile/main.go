package main


import (
	"fmt"
	"image"
  "github.com/pixelartexchange/artbase.server/pixelart"
)


func main() {

  fmt.Printf( "Hello, Pixel Art v%s!\n", pixelart.Version )

  path     := "../morepunks.png"
	tileSize := image.Point{24, 24}
  punks := pixelart.ReadImageComposite( path, &tileSize )

	fmt.Println( punks.Bounds() )
  //=> (0,0)-(600,960)

	punk := punks.Tile( 0 )
	fmt.Println( punk.Bounds() )
	//=>

	punk.Save( "./morepunk0.png" )


	punks.Tile( 18 ).Save( "./morepunk18.png" )
	punks.Tile( 40 ).Save( "./morepunk40.png" )
	punks.Tile( 88 ).Save( "./morepunk88.png" )

	punks.Tile( 0 ).Zoom( 4 ).Save( "./morepunk0@4x.png" )
	punks.Tile( 18 ).Zoom( 4 ).Save( "./morepunk18@4x.png" )
	punks.Tile( 40 ).Zoom( 4 ).Save( "./morepunk40@4x.png" )
	punks.Tile( 88 ).Zoom( 4 ).Save( "./morepunk88@4x.png" )


  fmt.Println( "Bye")
}