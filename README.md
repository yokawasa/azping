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

## Supplementary information

You can get a list of all supported regions for your current subscription with azure CLI:

> az account list-locations -o table
```
DisplayName           Latitude    Longitude    Name
--------------------  ----------  -----------  ------------------
East Asia             22.267      114.188      eastasia
Southeast Asia        1.283       103.833      southeastasia
Central US            41.5908     -93.6208     centralus
East US               37.3719     -79.8164     eastus
East US 2             36.6681     -78.3889     eastus2
West US               37.783      -122.417     westus
North Central US      41.8819     -87.6278     northcentralus
South Central US      29.4167     -98.5        southcentralus
North Europe          53.3478     -6.2597      northeurope
West Europe           52.3667     4.9          westeurope
Japan West            34.6939     135.5022     japanwest
Japan East            35.68       139.77       japaneast
Brazil South          -23.55      -46.633      brazilsouth
Australia East        -33.86      151.2094     australiaeast
Australia Southeast   -37.8136    144.9631     australiasoutheast
South India           12.9822     80.1636      southindia
Central India         18.5822     73.9197      centralindia
West India            19.088      72.868       westindia
Canada Central        43.653      -79.383      canadacentral
Canada East           46.817      -71.217      canadaeast
UK South              50.941      -0.799       uksouth
UK West               53.427      -3.084       ukwest
West Central US       40.890      -110.234     westcentralus
West US 2             47.233      -119.852     westus2
Korea Central         37.5665     126.9780     koreacentral
Korea South           35.1796     129.0756     koreasouth
France Central        46.3772     2.3730       francecentral
France South          43.8345     2.1972       francesouth
Australia Central     -35.3075    149.1244     australiacentral
Australia Central 2   -35.3075    149.1244     australiacentral2
UAE Central           24.466667   54.366669    uaecentral
UAE North             25.266666   55.316666    uaenorth
South Africa North    -25.731340  28.218370    southafricanorth
South Africa West     -34.075691  18.843266    southafricawest
Switzerland North     47.451542   8.564572     switzerlandnorth
Switzerland West      46.204391   6.143158     switzerlandwest
Germany North         53.073635   8.806422     germanynorth
Germany West Central  50.110924   8.682127     germanywestcentral
Norway West           58.969975   5.733107     norwaywest
Norway East           59.913868   10.752245    norwayeast
```

