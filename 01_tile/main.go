package main


import (
  "fmt"
  "github.com/learnpixelart/pixelart.go/pixelart"
)


func main() {
  fmt.Printf( "Hello, Pixel Art v%s!\n", pixelart.Version )

  path     := "../morepunks.png"
  tileSize := pixelart.Point{24, 24}
  punks    := pixelart.ReadImageComposite( path, &tileSize )

  fmt.Println( punks.Bounds() )
  //=> (0,0)-(600,960)

  punk := punks.Tile( 0 )
  fmt.Println( punk.Bounds() )
  //=> (0,0)-(24,24)

  punk.Save( "./morepunk0.png" )


  punks.Tile( 18 ).Save( "./morepunk18.png" )
  punks.Tile( 40 ).Save( "./morepunk40.png" )
  punks.Tile( 88 ).Save( "./morepunk88.png" )

  punks.Tile( 0 ).Zoom( 4 ).Save( "./morepunk0@4x.png" )
  punks.Tile( 18 ).Zoom( 4 ).Save( "./morepunk18@4x.png" )
  punks.Tile( 40 ).Zoom( 4 ).Save( "./morepunk40@4x.png" )
  punks.Tile( 88 ).Zoom( 4 ).Save( "./morepunk88@4x.png" )

  punks.Tile( 0 ).Background( "#638596" ).Zoom( 4 ).Save( "./morepunk0_(grayish)@4x.png" )
  punks.Tile( 18 ).Background( "#638596" ).Zoom( 4 ).Save( "./morepunk18_(grayish)@4x.png" )
  punks.Tile( 40 ).Background( "#638596" ).Zoom( 4 ).Save( "./morepunk40_(grayish)@4x.png" )
  punks.Tile( 88 ).Background( "#638596" ).Zoom( 4 ).Save( "./morepunk88_(grayish)@4x.png" )

  punks.Tile( 0 ).Background( "#638596" ).Mirror().Zoom( 4 ).Save( "./morephunk0_(grayish)@4x.png" )
  punks.Tile( 18 ).Background( "#638596" ).Mirror().Zoom( 4 ).Save( "./morephunk18_(grayish)@4x.png" )
  punks.Tile( 40 ).Background( "#638596" ).Mirror().Zoom( 4 ).Save( "./morephunk40_(grayish)@4x.png" )
  punks.Tile( 88 ).Background( "#638596" ).Mirror().Zoom( 4 ).Save( "./morephunk88_(grayish)@4x.png" )

  punks.Tile( 0 ).Ukraine().Zoom( 4 ).Save( "./morepunk0_flag(ukraine)@4x.png" )
  punks.Tile( 18 ).Ukraine().Zoom( 4 ).Save( "./morepunk18_flag(ukraine)@4x.png" )
  punks.Tile( 40 ).Ukraine().Zoom( 4 ).Save( "./morepunk40_flag(ukraine)@4x.png" )
  punks.Tile( 88 ).Ukraine().Zoom( 4 ).Save( "./morepunk88_flag(ukraine)@4x.png" )

	punks.Tile( 0 ).Silhouette("#ffdd00").Background("#0057b7").Zoom( 4 ).Save( "./morepunk0_silhouette(ukraine)@4x.png" )
	punks.Tile( 18 ).Silhouette("#0057b7").Background("#ffdd00").Zoom( 4 ).Save( "./morepunk18_silhouette(ukraine)@4x.png" )
	punks.Tile( 40 ).Silhouette("#ffdd00").Background("#0057b7").Zoom( 4 ).Save( "./morepunk40_silhouette(ukraine)@4x.png" )
	punks.Tile( 88 ).Silhouette("#0057b7").Background("#ffdd00").Zoom( 4 ).Save( "./morepunk88_silhouette(ukraine)@4x.png" )


  fmt.Println( "Bye")
}