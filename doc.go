/*
csvmask masks comma-separated values (CSV) data.
csvmask reads from the standard input and writes to the standard output.
csvmask expects the data utf-8 encoded.

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
		Returns string that masked n character to the right of the field value with c.
		The syntax is {{right n "c" .FieldN}}.
	left
		Returns string that masked n character to the left of the field value with c.
		The syntax is {{left n "c" .FieldN}}.

Examples

To mask second field of the record with hash:

	cat testdata/test.csv | csvmask -template "{{.Field1}},{{hash .Field2}},{{.Field3}}"

To skip header line:

	cat testdata/test.csv | csvmask -template "{{.Field1}},{{hash .Field2}},{{.Field3}}" -skipheader

If you want to mask data encoded other than utf-8, use tools like nkf:

	cat testdata/test-sjis.csv | nkf -w | csvmask -template "{{.Field1}},{{hash .Field2}},{{.Field3}}"

*/
package main
