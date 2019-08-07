// Copyright 2019 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Program azping pings Azure regions and reports about the latency.
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

// TODO(jbd): Add more zones.
var endpoints = map[string]string{
  "eastasia":             "azpingeastasia.blob.core.windows.net",
  "southeastasia":        "azpingsoutheastasia.blob.core.windows.net",
  "centralus":            "azpingcentralus.blob.core.windows.net",
  "eastus":               "azpingeastus.blob.core.windows.net",
  "eastus2":              "azpingeastus2.blob.core.windows.net",
  "westus":               "azpingwestus.blob.core.windows.net",
  "northcentralus":       "azpingnorthcentralus.blob.core.windows.net",
  "southcentralus":       "azpingsouthcentralus.blob.core.windows.net",
  "northeurope":          "azpingnortheurope.blob.core.windows.net",
  "westeurope":           "azpingwesteurope.blob.core.windows.net",
  "japanwest":            "azpingjapanwest.blob.core.windows.net",
  "japaneast":            "azpingjapaneast.blob.core.windows.net",
  "brazilsouth":          "azpingbrazilsouth.blob.core.windows.net",
  "australiaeast":        "azpingaustraliaeast.blob.core.windows.net",
  "australiasoutheast":   "azpingaustraliasoutheast.blob.core.windows.net",
  "southindia":           "azpingsouthindia.blob.core.windows.net",
  "centralindia":         "azpingcentralindia.blob.core.windows.net",
  "westindia":            "azpingwestindia.blob.core.windows.net",
  "canadacentral":        "azpingcanadacentral.blob.core.windows.net",
  "canadaeast":           "azpingcanadaeast.blob.core.windows.net",
  "uksouth":              "azpinguksouth.blob.core.windows.net",
  "ukwest":               "azpingukwest.blob.core.windows.net",
  "westcentralus":        "azpingwestcentralus.blob.core.windows.net",
  "westus2":              "azpingwestus2.blob.core.windows.net",
  "koreacentral":         "azpingkoreacentral.blob.core.windows.net",
  "koreasouth":           "azpingkoreasouth.blob.core.windows.net",
  "francecentral":        "azpingfrancecentral.blob.core.windows.net",
  "australiacentral":     "azpingaustraliacentral.blob.core.windows.net",
  "uaenorth":             "azpinguaenorth.blob.core.windows.net",
  "southafricanorth":     "azpingsouthafricanorth.blob.core.windows.net",
}

var (
	top         bool
	number      int // number of requests for each region
	concurrency int
	timeout     time.Duration
	csv         bool
	verbose     bool
	// TODO(jbd): Add payload options such as body size.

	client  *http.Client // TODO(jbd): One client per worker?
	inputs  chan input
	outputs chan output
)

func main() {
	flag.BoolVar(&top, "top", false, "")
	//flag.IntVar(&number, "n", 10, "")
	flag.IntVar(&number, "n", 5, "")
	flag.IntVar(&concurrency, "c", 10, "")
	flag.DurationVar(&timeout, "t", time.Duration(0), "")
	flag.BoolVar(&verbose, "v", false, "")
	flag.BoolVar(&csv, "csv", false, "")

	flag.Usage = usage
	flag.Parse()

	if number < 0 || concurrency <= 0 {
		usage()
	}
	if csv {
		verbose = false // if output is CSV, no need for verbose output
	}

	client = &http.Client{
		Timeout: timeout,
	}

	go start()
	inputs = make(chan input, concurrency)
	outputs = make(chan output, number*len(endpoints))
	for i := 0; i < number; i++ {
		for r, e := range endpoints {
			inputs <- input{region: r, endpoint: e}
		}
	}
	close(inputs)
	report()
}

func start() {
	for worker := 0; worker < concurrency; worker++ {
		go func() {
			for m := range inputs {
				m.HTTP()
			}
		}()
	}
}

func report() {
	m := make(map[string]output)
	for i := 0; i < number*len(endpoints); i++ {
		o := <-outputs

		a := m[o.region]

		a.region = o.region
		a.durations = append(a.durations, o.durations[0])
		a.errors += o.errors

		m[o.region] = a
	}
	all := make([]output, 0, len(m))
	for _, t := range m {
		all = append(all, t)
	}

	// sort all by median duration.
	sort.Slice(all, func(i, j int) bool {
		return all[i].median() < all[j].median()
	})

	if top {
		t := all[0].region
		if t == "global" {
			t = all[1].region
		}
		fmt.Print(t)
		return
	}

	tr := tabwriter.NewWriter(os.Stdout, 3, 2, 2, ' ', 0)
	for i, a := range all {
		fmt.Fprintf(tr, "%2d.\t[%v]\t%v", i+1, a.region, a.median())
		if a.errors > 0 {
			fmt.Fprintf(tr, "\t(%d errors)", a.errors)
		}
		fmt.Fprintln(tr)
	}
	tr.Flush()
}

func usage() {
	fmt.Println(usageText)
	os.Exit(0)
}

var usageText = `azping [options...]

Options:
-n   Number of requests to be made to each region.
     By default 5; can't be negative.
-c   Max number of requests to be made at any time.
     By default 10; can't be negative or zero.
-t   Timeout. By default, no timeout.
     Examples: "500ms", "1s", "1s500ms".
-top If true, only the top (non-global) region is printed.

-csv CSV output; disables verbose output.
-v   Verbose output.
`
