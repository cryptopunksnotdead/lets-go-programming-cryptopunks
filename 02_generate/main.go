package main


import (
  "fmt"
	"os"
	"log"
	"strings"
	"regexp"
	"encoding/csv"

	"github.com/learnpixelart/pixelart.go/pixelart"
)


func readCSV( path string ) [][]string {
  // open file
  f, err := os.Open( path )
  if err != nil {
		log.Fatal(err)
  }

  // remember to close the file at the end of the program
  defer f.Close()

  // read csv values using csv.Reader
  r := csv.NewReader(f)
  header, err := r.Read()   // "skip" first row with header
  if err != nil {
	  log.Fatal(err)
  }
  fmt.Println( header )

  data, err := r.ReadAll()
  if err != nil {
		log.Fatal(err)
  }

	fmt.Println( len(data) )
	//=> 100000

  return data
}



// allow (ignore):
//    space ( ),
//    underscore (_),
//    dash (-)
var normalizeRegexp = regexp.MustCompile( "[ _-]" )

func normalize( str string ) string {
    str = strings.ToLower( str )
		str = normalizeRegexp.ReplaceAllString( str, "" )
		return str
}



var dir = "../basic"

func generatePunk( values ...string ) *pixelart.Image {
  punkType       := values[0]
  attributeNames := values[1:len(values)]

  punkType = normalize( punkType )

	path   :=  dir + "/" + punkType + ".png"
	punk := pixelart.ReadImage( path )

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
		path       = dir + "/" + m_or_f + "/" + attributeName + ".png"
		attribute   := pixelart.ReadImage( path )

    punk.Paste( attribute )
	}

	return punk
}



func main() {
  fmt.Printf( "Hello, Pixel Art v%s!\n", pixelart.Version )

	// test drive
	// generate punk #0
	punk := generatePunk( "Female 2", "Earring", "Blonde Bob", "Green Eye Shadow" )
	punk.Save( "./punk0.png" )
	punk.Zoom(20).Save( "./punk0@20x.png" )

	// generate punk #1
	punk = generatePunk( "Male 1", "Smile", "Mohawk" )
	punk.Save( "./punk1.png" )
	punk.Zoom(20).Save( "./punk1@20x.png" )


	/////////////////////
	// try all 10 000 punks
  recs := readCSV( "../punks.csv" )
	fmt.Printf( "%d punk(s)\n", len( recs ) )
	//=> 10 000 punk(s)


	punk = generatePunk( recs[666]... )
	punk.Zoom( 4 ).Save( "./punk666@4x.png" )


	for i,rec := range recs {
		fmt.Printf( "==> %d - %v\n", i, rec )
    punk = generatePunk( rec... )

    name := fmt.Sprintf("punk%d", i )

    punk.Save( "./o/" + name + ".png" )
    punk.Zoom(20).Save( "./o/" + name + "@20x.png" )
 }


	punks := pixelart.NewImageComposite( 100, 100, &pixelart.Point{24, 24} )   // try 10x10 grid


	for i,rec := range recs {
		fmt.Printf( "==> %d - %v\n", i, rec )
    punk = generatePunk( rec... )
		punks.Add( punk )
	}

  punks.Save( "./o/punks.png" )


  fmt.Println( "Bye")
}
