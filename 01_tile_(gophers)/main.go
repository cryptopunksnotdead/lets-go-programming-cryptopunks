package main

import (
  "fmt"
  "github.com/pixelartexchange/artbase.server/pixelart"
)


func main() {

  fmt.Printf( "Hello, Pixel Art v%s!\n", pixelart.Version )

  path     := "../gophers.png"
  tileSize := pixelart.Point{32, 32}
  gophers  := pixelart.ReadImageComposite( path, &tileSize )

  fmt.Println( gophers.Bounds() )
  //=> (0,0)-(224,160)


	gopher := gophers.Tile( 0 )
  fmt.Println( gopher.Bounds() )
  //=> (0,0)-(32,32)

	gopher.Save( "./gopher0.png" )

  gophers.Tile( 1 ).Save( "./gopher1.png" )
  gophers.Tile( 16 ).Save( "./gopher16.png" )
  gophers.Tile( 28 ).Save( "./gopher28.png" )


	gophers.Tile( 0 ).Zoom( 4 ).Save( "./gopher0@4x.png" )
	gophers.Tile( 1 ).Zoom( 4 ).Save( "./gopher1@4x.png" )
  gophers.Tile( 16 ).Zoom( 4 ).Save( "./gopher16@4x.png" )
  gophers.Tile( 28 ).Zoom( 4 ).Save( "./gopher28@4x.png" )


	gophers.Tile( 0 ).Background( "#638596" ).Zoom( 4 ).Save( "./gopher0_(grayish)@4x.png" )
	gophers.Tile( 1 ).Background( "#638596" ).Zoom( 4 ).Save( "./gopher1_(grayish)@4x.png" )
  gophers.Tile( 16 ).Background( "#638596" ).Zoom( 4 ).Save( "./gopher16_(grayish)@4x.png" )
  gophers.Tile( 28 ).Background( "#638596" ).Zoom( 4 ).Save( "./gopher28_(grayish)@4x.png" )


	gophers.Tile( 0 ).Ukraine().Zoom( 4 ).Save( "./gopher0_flag(ukraine)@4x.png" )
	gophers.Tile( 1 ).Ukraine().Zoom( 4 ).Save( "./gopher1_flag(ukraine)@4x.png" )
  gophers.Tile( 16 ).Ukraine().Zoom( 4 ).Save( "./gopher16_flag(ukraine)@4x.png" )
  gophers.Tile( 28 ).Ukraine().Zoom( 4 ).Save( "./gopher28_flag(ukraine)@4x.png" )


	fmt.Println( "bye")
}