# punkverse Package

Let's You Generate (Pixel) Punk Images From Text Using Built-In Archetype and Attribute Images


Example:


``` go
package main


import (
  "github.com/cryptopunksnotdead/lets-go-programming-cryptopunks/punkverse"
)


func main() {
  punk := punkverse.Generate( "Female 2", "Earring", "Blonde Bob", "Green Eye Shadow" )
  punk.Save( "./punk0.png" )
  punk.Zoom( 10 ).Save( "./punk0@10x.png" )
  punk.Background( "#60A4F7" ).Mirror().Zoom( 10 ).Save( "./phunk0@10x.png" )

  punk = punkverse.Generate( "Male 1", "Smile", "Mohawk" )
  punk.Save( "./punk1.png" )
  punk.Zoom( 10 ).Save( "./punk1@10x.png" )
  punk.Background( "#60A4F7" ).Mirror().Zoom( 10 ).Save( "./phunk1@10x.png" )

  punk = punkverse.Generate( "Alien", "Headband" )
  punk.Save( "./alien.png" )
  punk.Zoom( 10 ).Save( "./alien@10x.png" )

  punk = punkverse.Generate( "Zombie", "Crazy Hair" )
  punk.Save( "./zombie.png" )
  punk.Zoom( 10 ).Save( "./zombie@10x.png" )


  ///
  //  try "raw" attribute images
  female2 := punkverse.Find( "Female 2" )
  female2.Save( "./female2.png" )
  female2.Zoom( 10 ).Save( "./female2@10x.png" )

  male1 := punkverse.Find( "Male 1" )
  male1.Save( "./male1.png" )
  male1.Zoom( 10 ).Save( "./male1@10x.png" )

  tophatM := punkverse.Find( "Top Hat", punkverse.M )
  tophatM.Save( "./tophat_(m).png" )
  tophatM.Zoom( 10 ).Save( "./tophat_(m)@10x.png" )

  crazyhairM := punkverse.Find( "Crazy Hair", punkverse.M )
  crazyhairM.Save( "./crazyhair_(m).png" )
  crazyhairM.Zoom( 10 ).Save( "./crazyhair_(m)@10x.png" )

  crazyhairF := punkverse.Find( "Crazy Hair", punkverse.F )
  crazyhairF.Save( "./crazyhair_(f).png" )
  crazyhairF.Zoom( 10 ).Save( "./crazyhair_(f)@10x.png" )
}
```


Or an example with more archetypes and attributes added:

``` go
package main


import (
  "fmt"
  "github.com/learnpixelart/pixelart.go/pixelart"
  "github.com/cryptopunksnotdead/lets-go-programming-cryptopunks/punkverse"
)


var dir = "../more"


func main() {
  fmt.Printf( "Hello, Pixel Art v%s!\n", pixelart.Version )

  ////
  // add more archetypes and attributes to built-in pixmap
  punkverse.Pixmap[ "demon" ]           = pixelart.ReadImage( dir + "/demon.png" )
  punkverse.Pixmap[ "birthdayhat_(m)" ] = pixelart.ReadImage( dir + "/m/birthdayhat.png" )
  punkverse.Pixmap[ "heartshades_(m)" ] = pixelart.ReadImage( dir + "/m/heartshades.png" )
  punkverse.Pixmap[ "lasereyes_(m)" ]   = pixelart.ReadImage( dir + "/m/lasereyes.png" )
  punkverse.Pixmap[ "birthdayhat_(f)" ] = pixelart.ReadImage( dir + "/f/birthdayhat.png" )
  punkverse.Pixmap[ "heartshades_(f)" ] = pixelart.ReadImage( dir + "/f/heartshades.png" )
  punkverse.Pixmap[ "lasereyes_(f)" ]   = pixelart.ReadImage( dir + "/f/lasereyes.png" )


  /////
  // try out new archetypes and attributes
  punk := punkverse.Generate( "Demon", "Smile", "Heart Shades" )
  punk.Save( "./demon.png" )
  punk.Zoom( 10 ).Save( "./demon@10x.png" )

  punk = punkverse.Generate( "Male 1", "Smile", "Birthday Hat", "Laser Eyes" )
  punk.Save( "./punk2.png" )
  punk.Zoom( 10 ).Save( "./punk2@10x.png" )

  punk = punkverse.Generate( "Female 2", "Birthday Hat", "Heart Shades" )
  punk.Save( "./punk3.png" )
  punk.Zoom( 10 ).Save( "./punk3@10x.png" )
}
```



## Questions? Comments?

Yes, you can. Post them on the [D.I.Y. Punk (Pixel) Art reddit](https://old.reddit.com/r/DIYPunkArt). Thanks.


