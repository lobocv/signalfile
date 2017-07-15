package main

import (
	"fmt"
)

type RadarFile struct {

	filepath string

	trace_headers []string
	file_headers []string


}

func newRadarFile(filepath string) *RadarFile {
	f := new(RadarFile)
	f.filepath = filepath
	return f
}


type SignalFile interface {


	define_file_header()
	define_trace_header()
	define_header()

	get_pre_header()
	get_header_definition()
	get_header()

	parse_header()
	render_header()

	format_header()

	open()
	close()
	read()
	write()
	copy()

}

func main () {
	f := newRadarFile("/tmp/radarfile1.rf")
	fmt.Printf(f.filepath)

}
