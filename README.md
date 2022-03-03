

# Let's Go! Programming Pixel Punk Profile Pictures & (Generative) Art with Go - Step-by-Step Book / Guide

_Inside Unique 24×24 Pixel Art on the Blockchain..._

by [Gerald Bauer](https://github.com/geraldb), et al




## Do-It-Yourself (DIY) - Yes, You Can! - Mint Your Own Punks in Original 24x24 Pixel Format or With 2X / 4X / 8X Zoom


### Step 0 -  Download a punks all-in-one composite image / collection


One time / first time only - Download
a punks pixel art collection
from the
[Awesome 24px Downloads - Free Pixel Art Collections in the 24x24 Format](https://github.com/cryptopunksnotdead/awesome-24px) page.

Let's try the 1000 More Punks
collection
in a single all-in-one 600×960px image (~180 kb) for free.
See [`morepunks.png` »](https://github.com/cryptopunksnotdead/awesome-24px/blob/master/collection/morepunks.png)




Let's create a program to mint (more) punk pixel art images.
Let's (re)use
the pixelart package
from the [artbase "right-clicker" zero-config web service / server](https://github.com/pixelartexchange/artbase.server).
To go get the package use:

```
$ go get github.com/pixelartexchange/artbase.server/pixelart
```




### Step 1 -  Read punk composite image


``` go
package main


import (
	"fmt"
	"image"
  "github.com/pixelartexchange/artbase.server/pixelart"
)


func main() {

  fmt.Printf( "Hello, Pixel Art v%s!\n", pixelart.Version )

  path     := "./morepunks.png"
	tileSize := image.Point{24, 24}
  punks := pixelart.ReadImageComposite( path, &tileSize )

	fmt.Println( punks.Bounds() )
  //=> (0,0)-(600,960)
}
```


### Step 2 - Start minting

Note: By default punks get saved in the original 24x24 pixel format
and the first punk starts at index zero, that is, `0`.
running up to 999.

Let's mint punk #0, #18, #40, and #88.
Add inside `func main()`:


``` go
punk := punks.Tile( 0 )
fmt.Println( punk.Bounds() )
//=>

punk.Save( "./morepunk0.png" )


punks.Tile( 18 ).Save( "./morepunk18.png" )
punks.Tile( 40 ).Save( "./morepunk40.png" )
punks.Tile( 88 ).Save( "./morepunk88.png" )
```

And voila!

![](i/morepunk0.png)
![](i/morepunk18.png)
![](i/morepunk40.png)
![](i/morepunk88.png)


Let's change the zoom factor:

``` go
punks.Tile( 0 ).Zoom( 4 ).Save( "./morepunk0@4x.png" )
punks.Tile( 18 ).Zoom( 4 ).Save( "./morepunk18@4x.png" )
punks.Tile( 40 ).Zoom( 4 ).Save( "./morepunk40@4x.png" )
punks.Tile( 88 ).Zoom( 4 ).Save( "./morepunk88@4x.png" )
```

And voila in 4x!

![](i/morepunk0@4x.png)
![](i/morepunk18@4x.png)
![](i/morepunk40@4x.png)
![](i/morepunk88@4x.png)



Proof-of the pudding.
If you want to run the ready-made sample
program in `01_tile/` yourself try:

```
$ cd 01_tile
$ go run main.go
```








