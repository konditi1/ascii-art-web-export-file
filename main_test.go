package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	webart "webart/handler"
)

var codes = []struct {
	name   string
	method string
}{
	{
		"test 1",
		"PUT",
	},
	{
		"test 2",
		"DELETE",
	},
	{
		"test 3",
		"PATCH",
	},
	{
		"test 4",
		"TRACE",
	},
}

var testCases = []struct {
	name                                                   string
	line1, line2, line3, line4, line5, line6, line7, line8 string
	form_data                                              string
}{
	{
		"Test1",
		" _              _   _          \n",
		"| |            | | | |         \n",
		"| |__     ___  | | | |   ___   \n",
		"|  _ \\   / _ \\ | | | |  / _ \\  \n",
		"| | | | |  __/ | | | | | (_) | \n",
		"|_| |_|  \\___| |_| |_|  \\___/  \n",
		"                               \n",
		"                               \n",
		"input=hello&banner=standard",
	},
	{
		"Test2",
		"                                 \n",
		"_|    _|          _| _|          \n",
		"_|    _|   _|_|   _| _|   _|_|   \n",
		"_|_|_|_| _|_|_|_| _| _| _|    _| \n",
		"_|    _| _|       _| _| _|    _| \n",
		"_|    _|   _|_|_| _| _|   _|_|   \n",
		"                                 \n",
		"                                 \n",
		"input=Hello&banner=shadow",
	},
	{
		"Test3",
		"                 \n",
		"o  o     o o     \n",
		"|  |     | |     \n",
		"O--O o-o | | o-o \n",
		"|  | |-&#39; | | | | \n",
		"o  o o-o o o o-o \n",
		"                 \n",
		"                 \n",
		"input=Hello&banner=thinkertoy",
	},
}

func TestStatusCode(t *testing.T) {
	for _, tc := range codes {
		t.Run(tc.name, func(t *testing.T) {
			wr := httptest.NewRecorder()
			req := httptest.NewRequest(tc.method, "/", nil)
			webart.Asciiweb(wr, req)
			if wr.Code != http.StatusMethodNotAllowed {
				t.Errorf("got HTTP status code %d , expected 405", wr.Code)
			}
		})
	}
}

func TestWeb_Ascii_Art(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.NewReader(tc.form_data)
			wr := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/ascii-art", input)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			webart.Asciiweb(wr, req)
			if wr.Code != http.StatusOK {
				t.Errorf("got HTTP status code %d , expected 200", wr.Code)
				return
			}

			expected := tc.line1 + tc.line2 + tc.line3 + tc.line4 + tc.line5 + tc.line6 + tc.line7 + tc.line8
			if !strings.Contains(wr.Body.String(), expected) {
				t.Errorf("Expected %v but got %v ", expected, wr.Body.String())
			}
		})
	}
}
