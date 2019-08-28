/*
csvmask masks comma-separated values (CSV) files.
csvmask reads from the standard input and writes to the standard output.

Usage:
	csvmask [flags]

The flags are:
	-help
		Show help.
	-skipheader
		Skip first line of file as header.
	-template string
		The template of output.
		It is same format as text/template's text.
	-version
		Show version.

Arguments

	.FieldN
		.FieldN referres each fields of the csv record.
		N is start from 1.

Functions

	hash
		Returns hash of the field value.
	checksum
		Returns checksum of the field value.
	right
		Returns string that masked n character to the right of the field value.
	left
		Returns string that masked n character to the left of the field value.

Examples

	TBD

*/
package main
