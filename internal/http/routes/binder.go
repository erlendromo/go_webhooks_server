package routes

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"webhooks/internal/constants"
)

type Binder struct {
	Endpoint string
	Params   map[string]interface{}
	ID       string
}

/*
Helper function to bind the request to the binding.Request struct
This is a dynamic function, meaning it will parse any and all query parameters and path parameters
*/
func BindRequest(r *http.Request) Binder {
	binder := Binder{
		Params: make(map[string]any, 0),
	}

	getPathFromRequest(&binder, strings.Split(r.URL.Path, constants.ROOT))
	parseQueryParams(binder, r.URL.Query())

	return binder
}

/*
This function will parse the path from the request,
and set the appropriate fields in the binding.Request struct
*/
func getPathFromRequest(binder *Binder, path []string) {
	url := constants.ROOT
	for _, p := range path {
		p = strings.TrimSpace(p)
		if p == constants.WEBHOOKS {
			url = fmt.Sprintf("%s%s%s", url, p, constants.ROOT)
		} else if len(p) > 0 {
			binder.ID = p
		}
	}
	binder.Endpoint = url
}

/*
This function will parse the query parameters from the request,
and set the appropriate fields in the binding.Request struct dynamically
*/
func parseQueryParams(binder Binder, queryMap url.Values) {
	for key, value := range queryMap {
		valueSlice := make([]any, 0)
		if len(value) < 1 {
			continue
		} else if len(value) == 1 {
			parseValue := getValueType(value[0])
			binder.Params[key] = parseValue
		} else {
			for _, v := range value {
				parseValue := getValueType(v)
				valueSlice = append(valueSlice, parseValue)
			}
			binder.Params[key] = valueSlice
		}
	}
}

/*
This function will parse the value type from the query parameters
(Only supports float64, int, bool, and string at the moment - can be extended to support more types)
*/
func getValueType(value string) any {
	vFloat, err := strconv.ParseFloat(value, 64)
	if err == nil {
		return vFloat
	}

	vInt, err := strconv.Atoi(value)
	if err == nil {
		return vInt
	}

	vBool, err := strconv.ParseBool(value)
	if err == nil {
		return vBool
	}

	return value
}
