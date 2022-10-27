package punkverse

import (
  // "fmt"
	"strings"
	"regexp"
	"log"

	"github.com/learnpixelart/pixelart.go/pixelart"
)



// allow (ignore):
//    space ( ),
//    underscore (_),
//    dash (-)
var normalizeRegexp = regexp.MustCompile( "[ \t_-]" )

func normalize( str string ) string {
    str = strings.ToLower( str )
		str = normalizeRegexp.ReplaceAllString( str, "" )
		return str
}


/////
//  todo/check:
//     find/use a different (more idomatic?) interface to lookup attribute (images)
//      - why? why not?
//
//    punkverse.Find( "Alien" )
//
//    punkverse.Find( "Headband", punkverse.MALE )
//    punkverse.Find( "Headband", punkverse.M )
//    punkverse.Find( "Headband", "m" )
//
//    punkverse.Find( "Headband", punkverse.FEMALE )
//    punkverse.Find( "Headband", punkverse.F )
//    punkverse.Find( "Headband", "f" )

const (
	 MALE   = "m"
	 FEMALE = "f"
	 M      = "m"
	 F      = "f"
)

func Find( q string, opts... string  ) *pixelart.Image {
   key := normalize( q )

   for i, opt := range opts {
     if i > 0 {
			 // note: only zero or one option allowed for now; sorry
			 log.Fatalf( "[punkverse.Find] ERROR - only zero or one option allowed for punkverse.Find; got >%v<", opts )
		 }

		 if opt == "m" {
			 key = key + "_(m)"
		 } else if opt == "f" {
			 key = key + "_(f)"
		 } else {
			 log.Fatalf( "[punkverse.Find] ERROR - unknown option >%v<for punkverse.Find", opt )
		 }
	 }

	 if pix, ok := Pixmap[ key ]; ok {
      return pix
	 } else {
      // panic - attribute not found !!
			log.Fatalf( "[punkverse.Find] ERROR - attribute not found for >%s< using lookup key >%s<", q, key )
			return nil
	 }
}


func Generate( values ...string ) *pixelart.Image {
  punkType       := values[0]
  attributeNames := values[1:len(values)]

	canvas := pixelart.NewImage( 24, 24 )

  punkType = normalize( punkType )
	punk := Find( punkType )
  canvas.Paste( punk )

	var m_or_f string
	if strings.Index( punkType, "female" ) != -1 {
		m_or_f = "f"
	} else {
		m_or_f = "m"
	}


	for _, attributeName := range attributeNames {
		if attributeName == "" {   // skip empty attributes
			continue
		}

		attributeName = normalize( attributeName )
		attribute   := Find( attributeName, m_or_f )

    canvas.Paste( attribute )
	}

	return canvas
}




/////
// (more) convenience helper

func makeImage24x24( pix []uint32 ) *pixelart.Image {
  return pixelart.MakeImage( pix, &pixelart.Point{ 24,24 })
}
