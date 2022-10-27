//////
//  "script" to export builtin archtetype and attributes for punkverse

package main


import (
  "fmt"
	"bytes"
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




var dir = "../basic"


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



func main() {
  fmt.Printf( "Hello, Pixel Art v%s!\n", pixelart.Version )

	recs := readCSV( "./attributes.csv" )
	fmt.Printf( "  %d attribute(s)\n", len(recs))
	// 133 attribute(s)


	var buf1 bytes.Buffer   // all pixmaps
	buf1.WriteString( "package punkverse\n" )
	buf1.WriteString( "\n\n")


  var buf2 bytes.Buffer   // all pixmap mappings
	buf2.WriteString( `package punkverse

import (
		"github.com/pixelartexchange/artbase.server/pixelart"
)` )
  buf2.WriteString( "\n\n" )
  buf2.WriteString( "var Pixmap = map[string]*pixelart.Image{\n" )

	// export pixels
	for i,rec := range recs {
		path := dir + "/" + rec[0]
		fmt.Printf( "==> [%d] %s reading...\n", i, path )

		attribute := pixelart.ReadImage( path )
		fmt.Println( attribute.Bounds() )

    pix := attribute.Export()

		key := normalize( rec[1] )
    m_or_f := strings.TrimSpace( rec[2] )

		if m_or_f == "m" {
			key = key + "_(m)"
		}
		if m_or_f == "f" {
			key = key + "_(f)"
		}

	  name := normalize( rec[1] )
		if m_or_f == "m" {
			name = name + "M"
		}
		if m_or_f == "f" {
			name = name + "F"
		}

		name = strings.ToUpper( name[0:1] ) + name[1:]

		buf1.WriteString( fmt.Sprintf( "var pix%s = []uint32{\n", name ))
		buf1.WriteString( pix )
		buf1.WriteString( "}\n" )
		buf1.WriteString( "\n\n" )

		buf2.WriteString( fmt.Sprintf( "  \"%s\": makeImage24x24( pix%s ),\n", key, name ))
	}

  buf2.WriteString( "}\n" )


 	// save bufs
	 os.WriteFile( "./o/pix.go", buf1.Bytes(), 0644 )
	 os.WriteFile( "./o/pixmap.go", buf2.Bytes(), 0644 )


	fmt.Println( "bye")
}

