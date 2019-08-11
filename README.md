# azping

azping is a command line tools that reports median latency to
Azure regions. It is a fork of [gcping](https://github.com/GoogleCloudPlatform/gcping).

```
azping [options...]

Options:
-n   Number of requests to be made to each region.
     By default 5; can't be negative.
-c   Max number of requests to be made at any time.
     By default 10; can't be negative or zero.
-t   Timeout. By default, no timeout.
     Examples: "500ms", "1s", "1s500ms".
-top If true, only the top region is printed.

-csv CSV output; disables verbose output.
-v   Verbose output.
```

An example output:

```
$ azping

 1.  [japaneast]           44.76367ms
 2.  [japanwest]           51.818528ms
 3.  [koreacentral]        82.854564ms
 4.  [koreasouth]          106.661344ms
 5.  [eastasia]            161.402128ms
 6.  [southeastasia]       191.654551ms
 7.  [australiaeast]       239.64985ms
 8.  [westus2]             241.140593ms
 9.  [westus]              247.800051ms
10.  [southindia]          260.301905ms
11.  [australiacentral]    264.562011ms
12.  [australiasoutheast]  294.994037ms
13.  [centralus]           302.262868ms
14.  [southcentralus]      304.422153ms
15.  [centralindia]        306.603511ms
16.  [westindia]           308.15552ms
17.  [westcentralus]       320.195892ms
18.  [uaenorth]            341.351333ms
19.  [canadacentral]       344.130208ms
20.  [canadaeast]          362.906659ms
21.  [eastus2]             367.944441ms
22.  [eastus]              385.297558ms
23.  [northeurope]         487.650428ms
24.  [uksouth]             504.534594ms
25.  [francecentral]       518.567778ms
26.  [ukwest]              521.46691ms
27.  [westeurope]          563.110408ms
28.  [brazilsouth]         563.938323ms
29.  [northcentralus]      597.346426ms
30.  [southafricanorth]    803.077525ms
```

## Installation

* Linux 64-bit: https://azpingrelease.blob.core.windows.net/azping_linux_amd64
  ```
  $ curl https://azpingrelease.blob.core.windows.net/azping_linux_amd64 > azping && chmod +x azping
  ```
* Mac 64-bit: https://azpingrelease.blob.core.windows.net/azping_darwin_amd64
* Windows 64-bit: https://azpingrelease.blob.core.windows.net/azping_windows_amd64

Or, you can always build the binary from the source code like this:

```
$ git clone https://github.com/yokawasa/azping.git
$ cd azping
$ make
$ tree bin

bin
├── azping_darwin_amd64
├── azping_linux_amd64
└── azping_windows_amd64
```
