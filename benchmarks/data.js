window.BENCHMARK_DATA = {
  "lastUpdate": 1784479289033,
  "repoUrl": "https://github.com/gofiber/schema",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "email": "rene@gofiber.io",
            "name": "René",
            "username": "ReneWerner87"
          },
          "committer": {
            "email": "rene@gofiber.io",
            "name": "René",
            "username": "ReneWerner87"
          },
          "distinct": true,
          "id": "706e5d730ac6333129f2e185cfef23b77b925b41",
          "message": "feat(benchmark): integrate reusable benchmark workflow and update README with benchmark links",
          "timestamp": "2026-07-11T23:54:16+02:00",
          "tree_id": "182b8161436765070ec8087fad0a146096cc0083",
          "url": "https://github.com/gofiber/schema/commit/706e5d730ac6333129f2e185cfef23b77b925b41"
        },
        "date": 1783807046698,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBool",
            "value": 3.528,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "339153373 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBool - ns/op",
            "value": 3.528,
            "unit": "ns/op",
            "extra": "339153373 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBool - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "339153373 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBool - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "339153373 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertInt",
            "value": 7.406,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "161983354 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertInt - ns/op",
            "value": 7.406,
            "unit": "ns/op",
            "extra": "161983354 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertInt - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "161983354 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertInt - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "161983354 times\n4 procs"
          },
          {
            "name": "BenchmarkAll",
            "value": 32.51,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "37147578 times\n4 procs"
          },
          {
            "name": "BenchmarkAll - ns/op",
            "value": 32.51,
            "unit": "ns/op",
            "extra": "37147578 times\n4 procs"
          },
          {
            "name": "BenchmarkAll - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "37147578 times\n4 procs"
          },
          {
            "name": "BenchmarkAll - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "37147578 times\n4 procs"
          },
          {
            "name": "BenchmarkDecoderMultipartFiles",
            "value": 6025,
            "unit": "ns/op\t    1567 B/op\t      48 allocs/op",
            "extra": "192255 times\n4 procs"
          },
          {
            "name": "BenchmarkDecoderMultipartFiles - ns/op",
            "value": 6025,
            "unit": "ns/op",
            "extra": "192255 times\n4 procs"
          },
          {
            "name": "BenchmarkDecoderMultipartFiles - B/op",
            "value": 1567,
            "unit": "B/op",
            "extra": "192255 times\n4 procs"
          },
          {
            "name": "BenchmarkDecoderMultipartFiles - allocs/op",
            "value": 48,
            "unit": "allocs/op",
            "extra": "192255 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-0",
            "value": 7.778,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "154048674 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-0 - ns/op",
            "value": 7.778,
            "unit": "ns/op",
            "extra": "154048674 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-0 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "154048674 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-0 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "154048674 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-1",
            "value": 7.751,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "154856840 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-1 - ns/op",
            "value": 7.751,
            "unit": "ns/op",
            "extra": "154856840 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-1 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "154856840 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-1 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "154856840 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-2",
            "value": 6.65,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "181536154 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-2 - ns/op",
            "value": 6.65,
            "unit": "ns/op",
            "extra": "181536154 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-2 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "181536154 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-2 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "181536154 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-3",
            "value": 7.752,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "154739546 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-3 - ns/op",
            "value": 7.752,
            "unit": "ns/op",
            "extra": "154739546 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-3 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "154739546 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-3 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "154739546 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-4",
            "value": 4.241,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "283878244 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-4 - ns/op",
            "value": 4.241,
            "unit": "ns/op",
            "extra": "283878244 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-4 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "283878244 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-4 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "283878244 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-5",
            "value": 7.764,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "154846917 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-5 - ns/op",
            "value": 7.764,
            "unit": "ns/op",
            "extra": "154846917 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-5 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "154846917 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-5 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "154846917 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-6",
            "value": 13.75,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "91857884 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-6 - ns/op",
            "value": 13.75,
            "unit": "ns/op",
            "extra": "91857884 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-6 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "91857884 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-6 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "91857884 times\n4 procs"
          },
          {
            "name": "BenchmarkHandleMultipartField",
            "value": 176.4,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "6835897 times\n4 procs"
          },
          {
            "name": "BenchmarkHandleMultipartField - ns/op",
            "value": 176.4,
            "unit": "ns/op",
            "extra": "6835897 times\n4 procs"
          },
          {
            "name": "BenchmarkHandleMultipartField - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "6835897 times\n4 procs"
          },
          {
            "name": "BenchmarkHandleMultipartField - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "6835897 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecode",
            "value": 6568,
            "unit": "ns/op\t    1200 B/op\t      62 allocs/op",
            "extra": "180806 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecode - ns/op",
            "value": 6568,
            "unit": "ns/op",
            "extra": "180806 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecode - B/op",
            "value": 1200,
            "unit": "B/op",
            "extra": "180806 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecode - allocs/op",
            "value": 62,
            "unit": "allocs/op",
            "extra": "180806 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecodeParallel",
            "value": 3247,
            "unit": "ns/op\t    1200 B/op\t      62 allocs/op",
            "extra": "359539 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecodeParallel - ns/op",
            "value": 3247,
            "unit": "ns/op",
            "extra": "359539 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecodeParallel - B/op",
            "value": 1200,
            "unit": "B/op",
            "extra": "359539 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecodeParallel - allocs/op",
            "value": 62,
            "unit": "allocs/op",
            "extra": "359539 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructDecode",
            "value": 2912,
            "unit": "ns/op\t     479 B/op\t      30 allocs/op",
            "extra": "376902 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructDecode - ns/op",
            "value": 2912,
            "unit": "ns/op",
            "extra": "376902 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructDecode - B/op",
            "value": 479,
            "unit": "B/op",
            "extra": "376902 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructDecode - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "376902 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckRequiredFields",
            "value": 1148,
            "unit": "ns/op\t    1048 B/op\t      15 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckRequiredFields - ns/op",
            "value": 1148,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckRequiredFields - B/op",
            "value": 1048,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckRequiredFields - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationDecoding",
            "value": 668.9,
            "unit": "ns/op\t     184 B/op\t       8 allocs/op",
            "extra": "1788633 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationDecoding - ns/op",
            "value": 668.9,
            "unit": "ns/op",
            "extra": "1788633 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationDecoding - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "1788633 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationDecoding - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1788633 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncode",
            "value": 1589,
            "unit": "ns/op\t     534 B/op\t       5 allocs/op",
            "extra": "846962 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncode - ns/op",
            "value": 1589,
            "unit": "ns/op",
            "extra": "846962 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncode - B/op",
            "value": 534,
            "unit": "B/op",
            "extra": "846962 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncode - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "846962 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncodeParallel",
            "value": 714.6,
            "unit": "ns/op\t     585 B/op\t       5 allocs/op",
            "extra": "1725594 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncodeParallel - ns/op",
            "value": 714.6,
            "unit": "ns/op",
            "extra": "1725594 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncodeParallel - B/op",
            "value": 585,
            "unit": "B/op",
            "extra": "1725594 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncodeParallel - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "1725594 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncode",
            "value": 2736,
            "unit": "ns/op\t     998 B/op\t      12 allocs/op",
            "extra": "446582 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncode - ns/op",
            "value": 2736,
            "unit": "ns/op",
            "extra": "446582 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncode - B/op",
            "value": 998,
            "unit": "B/op",
            "extra": "446582 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncode - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "446582 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncodeParallel",
            "value": 1297,
            "unit": "ns/op\t    1082 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncodeParallel - ns/op",
            "value": 1297,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncodeParallel - B/op",
            "value": 1082,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncodeParallel - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationEncoding",
            "value": 294.3,
            "unit": "ns/op\t     136 B/op\t       2 allocs/op",
            "extra": "3887112 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationEncoding - ns/op",
            "value": 294.3,
            "unit": "ns/op",
            "extra": "3887112 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationEncoding - B/op",
            "value": 136,
            "unit": "B/op",
            "extra": "3887112 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationEncoding - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "3887112 times\n4 procs"
          },
          {
            "name": "BenchmarkMultiErrorError",
            "value": 167.5,
            "unit": "ns/op\t      40 B/op\t       2 allocs/op",
            "extra": "7163599 times\n4 procs"
          },
          {
            "name": "BenchmarkMultiErrorError - ns/op",
            "value": 167.5,
            "unit": "ns/op",
            "extra": "7163599 times\n4 procs"
          },
          {
            "name": "BenchmarkMultiErrorError - B/op",
            "value": 40,
            "unit": "B/op",
            "extra": "7163599 times\n4 procs"
          },
          {
            "name": "BenchmarkMultiErrorError - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "7163599 times\n4 procs"
          },
          {
            "name": "BenchmarkParseTag",
            "value": 49.59,
            "unit": "ns/op\t      32 B/op\t       1 allocs/op",
            "extra": "24163434 times\n4 procs"
          },
          {
            "name": "BenchmarkParseTag - ns/op",
            "value": 49.59,
            "unit": "ns/op",
            "extra": "24163434 times\n4 procs"
          },
          {
            "name": "BenchmarkParseTag - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "24163434 times\n4 procs"
          },
          {
            "name": "BenchmarkParseTag - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "24163434 times\n4 procs"
          },
          {
            "name": "BenchmarkIsZero",
            "value": 18.8,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "61707901 times\n4 procs"
          },
          {
            "name": "BenchmarkIsZero - ns/op",
            "value": 18.8,
            "unit": "ns/op",
            "extra": "61707901 times\n4 procs"
          },
          {
            "name": "BenchmarkIsZero - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "61707901 times\n4 procs"
          },
          {
            "name": "BenchmarkIsZero - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "61707901 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertPointer",
            "value": 22.8,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "49980931 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertPointer - ns/op",
            "value": 22.8,
            "unit": "ns/op",
            "extra": "49980931 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertPointer - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "49980931 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertPointer - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "49980931 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "rene@gofiber.io",
            "name": "RW",
            "username": "ReneWerner87"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "36dffb66bdfa355cddc97270426a4c50e32d7363",
          "message": "Merge pull request #72 from gofiber/dependabot/go_modules/github.com/gofiber/utils/v2-2.1.2\n\nbuild(deps): bump github.com/gofiber/utils/v2 from 2.1.1 to 2.1.2",
          "timestamp": "2026-07-12T18:51:34+02:00",
          "tree_id": "c1190a726bf598dbf11621eced7a97eafe179b19",
          "url": "https://github.com/gofiber/schema/commit/36dffb66bdfa355cddc97270426a4c50e32d7363"
        },
        "date": 1783875235722,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBool",
            "value": 3.441,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "349000953 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBool - ns/op",
            "value": 3.441,
            "unit": "ns/op",
            "extra": "349000953 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBool - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "349000953 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBool - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "349000953 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertInt",
            "value": 8.426,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "142413852 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertInt - ns/op",
            "value": 8.426,
            "unit": "ns/op",
            "extra": "142413852 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertInt - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "142413852 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertInt - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "142413852 times\n4 procs"
          },
          {
            "name": "BenchmarkAll",
            "value": 34.85,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "34542174 times\n4 procs"
          },
          {
            "name": "BenchmarkAll - ns/op",
            "value": 34.85,
            "unit": "ns/op",
            "extra": "34542174 times\n4 procs"
          },
          {
            "name": "BenchmarkAll - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "34542174 times\n4 procs"
          },
          {
            "name": "BenchmarkAll - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "34542174 times\n4 procs"
          },
          {
            "name": "BenchmarkDecoderMultipartFiles",
            "value": 6239,
            "unit": "ns/op\t    1567 B/op\t      48 allocs/op",
            "extra": "195333 times\n4 procs"
          },
          {
            "name": "BenchmarkDecoderMultipartFiles - ns/op",
            "value": 6239,
            "unit": "ns/op",
            "extra": "195333 times\n4 procs"
          },
          {
            "name": "BenchmarkDecoderMultipartFiles - B/op",
            "value": 1567,
            "unit": "B/op",
            "extra": "195333 times\n4 procs"
          },
          {
            "name": "BenchmarkDecoderMultipartFiles - allocs/op",
            "value": 48,
            "unit": "allocs/op",
            "extra": "195333 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-0",
            "value": 7.176,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "167185224 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-0 - ns/op",
            "value": 7.176,
            "unit": "ns/op",
            "extra": "167185224 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-0 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "167185224 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-0 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "167185224 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-1",
            "value": 7.186,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "166849693 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-1 - ns/op",
            "value": 7.186,
            "unit": "ns/op",
            "extra": "166849693 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-1 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "166849693 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-1 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "166849693 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-2",
            "value": 6.053,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "201795003 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-2 - ns/op",
            "value": 6.053,
            "unit": "ns/op",
            "extra": "201795003 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-2 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "201795003 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-2 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "201795003 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-3",
            "value": 7.167,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "167408941 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-3 - ns/op",
            "value": 7.167,
            "unit": "ns/op",
            "extra": "167408941 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-3 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "167408941 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-3 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "167408941 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-4",
            "value": 4.051,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "295939225 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-4 - ns/op",
            "value": 4.051,
            "unit": "ns/op",
            "extra": "295939225 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-4 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "295939225 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-4 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "295939225 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-5",
            "value": 7.175,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "167215017 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-5 - ns/op",
            "value": 7.175,
            "unit": "ns/op",
            "extra": "167215017 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-5 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "167215017 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-5 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "167215017 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-6",
            "value": 12.95,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "90965682 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-6 - ns/op",
            "value": 12.95,
            "unit": "ns/op",
            "extra": "90965682 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-6 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "90965682 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-6 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "90965682 times\n4 procs"
          },
          {
            "name": "BenchmarkHandleMultipartField",
            "value": 183.7,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "6493064 times\n4 procs"
          },
          {
            "name": "BenchmarkHandleMultipartField - ns/op",
            "value": 183.7,
            "unit": "ns/op",
            "extra": "6493064 times\n4 procs"
          },
          {
            "name": "BenchmarkHandleMultipartField - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "6493064 times\n4 procs"
          },
          {
            "name": "BenchmarkHandleMultipartField - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "6493064 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecode",
            "value": 6935,
            "unit": "ns/op\t    1200 B/op\t      62 allocs/op",
            "extra": "169494 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecode - ns/op",
            "value": 6935,
            "unit": "ns/op",
            "extra": "169494 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecode - B/op",
            "value": 1200,
            "unit": "B/op",
            "extra": "169494 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecode - allocs/op",
            "value": 62,
            "unit": "allocs/op",
            "extra": "169494 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecodeParallel",
            "value": 3406,
            "unit": "ns/op\t    1200 B/op\t      62 allocs/op",
            "extra": "345160 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecodeParallel - ns/op",
            "value": 3406,
            "unit": "ns/op",
            "extra": "345160 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecodeParallel - B/op",
            "value": 1200,
            "unit": "B/op",
            "extra": "345160 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecodeParallel - allocs/op",
            "value": 62,
            "unit": "allocs/op",
            "extra": "345160 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructDecode",
            "value": 3092,
            "unit": "ns/op\t     479 B/op\t      30 allocs/op",
            "extra": "373644 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructDecode - ns/op",
            "value": 3092,
            "unit": "ns/op",
            "extra": "373644 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructDecode - B/op",
            "value": 479,
            "unit": "B/op",
            "extra": "373644 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructDecode - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "373644 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckRequiredFields",
            "value": 1182,
            "unit": "ns/op\t    1048 B/op\t      15 allocs/op",
            "extra": "972890 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckRequiredFields - ns/op",
            "value": 1182,
            "unit": "ns/op",
            "extra": "972890 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckRequiredFields - B/op",
            "value": 1048,
            "unit": "B/op",
            "extra": "972890 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckRequiredFields - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "972890 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationDecoding",
            "value": 782.6,
            "unit": "ns/op\t     184 B/op\t       8 allocs/op",
            "extra": "1526161 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationDecoding - ns/op",
            "value": 782.6,
            "unit": "ns/op",
            "extra": "1526161 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationDecoding - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "1526161 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationDecoding - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1526161 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncode",
            "value": 1628,
            "unit": "ns/op\t     590 B/op\t       5 allocs/op",
            "extra": "746356 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncode - ns/op",
            "value": 1628,
            "unit": "ns/op",
            "extra": "746356 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncode - B/op",
            "value": 590,
            "unit": "B/op",
            "extra": "746356 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncode - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "746356 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncodeParallel",
            "value": 747.5,
            "unit": "ns/op\t     558 B/op\t       5 allocs/op",
            "extra": "1625156 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncodeParallel - ns/op",
            "value": 747.5,
            "unit": "ns/op",
            "extra": "1625156 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncodeParallel - B/op",
            "value": 558,
            "unit": "B/op",
            "extra": "1625156 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncodeParallel - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "1625156 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncode",
            "value": 2654,
            "unit": "ns/op\t    1014 B/op\t      12 allocs/op",
            "extra": "435298 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncode - ns/op",
            "value": 2654,
            "unit": "ns/op",
            "extra": "435298 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncode - B/op",
            "value": 1014,
            "unit": "B/op",
            "extra": "435298 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncode - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "435298 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncodeParallel",
            "value": 1440,
            "unit": "ns/op\t    1100 B/op\t      12 allocs/op",
            "extra": "925639 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncodeParallel - ns/op",
            "value": 1440,
            "unit": "ns/op",
            "extra": "925639 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncodeParallel - B/op",
            "value": 1100,
            "unit": "B/op",
            "extra": "925639 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncodeParallel - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "925639 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationEncoding",
            "value": 312.8,
            "unit": "ns/op\t     141 B/op\t       2 allocs/op",
            "extra": "3700627 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationEncoding - ns/op",
            "value": 312.8,
            "unit": "ns/op",
            "extra": "3700627 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationEncoding - B/op",
            "value": 141,
            "unit": "B/op",
            "extra": "3700627 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationEncoding - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "3700627 times\n4 procs"
          },
          {
            "name": "BenchmarkMultiErrorError",
            "value": 176.2,
            "unit": "ns/op\t      40 B/op\t       2 allocs/op",
            "extra": "6762783 times\n4 procs"
          },
          {
            "name": "BenchmarkMultiErrorError - ns/op",
            "value": 176.2,
            "unit": "ns/op",
            "extra": "6762783 times\n4 procs"
          },
          {
            "name": "BenchmarkMultiErrorError - B/op",
            "value": 40,
            "unit": "B/op",
            "extra": "6762783 times\n4 procs"
          },
          {
            "name": "BenchmarkMultiErrorError - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "6762783 times\n4 procs"
          },
          {
            "name": "BenchmarkParseTag",
            "value": 47.89,
            "unit": "ns/op\t      32 B/op\t       1 allocs/op",
            "extra": "24125967 times\n4 procs"
          },
          {
            "name": "BenchmarkParseTag - ns/op",
            "value": 47.89,
            "unit": "ns/op",
            "extra": "24125967 times\n4 procs"
          },
          {
            "name": "BenchmarkParseTag - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "24125967 times\n4 procs"
          },
          {
            "name": "BenchmarkParseTag - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "24125967 times\n4 procs"
          },
          {
            "name": "BenchmarkIsZero",
            "value": 18.99,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "62124363 times\n4 procs"
          },
          {
            "name": "BenchmarkIsZero - ns/op",
            "value": 18.99,
            "unit": "ns/op",
            "extra": "62124363 times\n4 procs"
          },
          {
            "name": "BenchmarkIsZero - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "62124363 times\n4 procs"
          },
          {
            "name": "BenchmarkIsZero - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "62124363 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertPointer",
            "value": 23.16,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "49450338 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertPointer - ns/op",
            "value": 23.16,
            "unit": "ns/op",
            "extra": "49450338 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertPointer - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "49450338 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertPointer - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "49450338 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "rene@gofiber.io",
            "name": "RW",
            "username": "ReneWerner87"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "860b052580e4445400f6fc5f578f7f9ee47ff82d",
          "message": "Merge pull request #73 from gofiber/fix/pathcache-detach-key\n\nfix(cache): clone path key before storing in pathCache",
          "timestamp": "2026-07-13T12:51:01+02:00",
          "tree_id": "96fa54b4e92eb620e34107726c766ad8a6ed23e2",
          "url": "https://github.com/gofiber/schema/commit/860b052580e4445400f6fc5f578f7f9ee47ff82d"
        },
        "date": 1783939909943,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBool",
            "value": 3.434,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "347948488 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBool - ns/op",
            "value": 3.434,
            "unit": "ns/op",
            "extra": "347948488 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBool - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "347948488 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBool - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "347948488 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertInt",
            "value": 8.453,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "141878338 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertInt - ns/op",
            "value": 8.453,
            "unit": "ns/op",
            "extra": "141878338 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertInt - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "141878338 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertInt - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "141878338 times\n4 procs"
          },
          {
            "name": "BenchmarkAll",
            "value": 33.91,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "35128838 times\n4 procs"
          },
          {
            "name": "BenchmarkAll - ns/op",
            "value": 33.91,
            "unit": "ns/op",
            "extra": "35128838 times\n4 procs"
          },
          {
            "name": "BenchmarkAll - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "35128838 times\n4 procs"
          },
          {
            "name": "BenchmarkAll - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "35128838 times\n4 procs"
          },
          {
            "name": "BenchmarkDecoderMultipartFiles",
            "value": 6151,
            "unit": "ns/op\t    1567 B/op\t      48 allocs/op",
            "extra": "189445 times\n4 procs"
          },
          {
            "name": "BenchmarkDecoderMultipartFiles - ns/op",
            "value": 6151,
            "unit": "ns/op",
            "extra": "189445 times\n4 procs"
          },
          {
            "name": "BenchmarkDecoderMultipartFiles - B/op",
            "value": 1567,
            "unit": "B/op",
            "extra": "189445 times\n4 procs"
          },
          {
            "name": "BenchmarkDecoderMultipartFiles - allocs/op",
            "value": 48,
            "unit": "allocs/op",
            "extra": "189445 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-0",
            "value": 7.161,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "166910278 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-0 - ns/op",
            "value": 7.161,
            "unit": "ns/op",
            "extra": "166910278 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-0 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "166910278 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-0 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "166910278 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-1",
            "value": 6.87,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "174728102 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-1 - ns/op",
            "value": 6.87,
            "unit": "ns/op",
            "extra": "174728102 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-1 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "174728102 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-1 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "174728102 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-2",
            "value": 5.613,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "213424458 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-2 - ns/op",
            "value": 5.613,
            "unit": "ns/op",
            "extra": "213424458 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-2 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "213424458 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-2 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "213424458 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-3",
            "value": 6.867,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "174704047 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-3 - ns/op",
            "value": 6.867,
            "unit": "ns/op",
            "extra": "174704047 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-3 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "174704047 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-3 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "174704047 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-4",
            "value": 3.74,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "320556444 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-4 - ns/op",
            "value": 3.74,
            "unit": "ns/op",
            "extra": "320556444 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-4 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "320556444 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-4 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "320556444 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-5",
            "value": 6.856,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "174995996 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-5 - ns/op",
            "value": 6.856,
            "unit": "ns/op",
            "extra": "174995996 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-5 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "174995996 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-5 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "174995996 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-6",
            "value": 12.78,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "95706830 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-6 - ns/op",
            "value": 12.78,
            "unit": "ns/op",
            "extra": "95706830 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-6 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "95706830 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-6 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "95706830 times\n4 procs"
          },
          {
            "name": "BenchmarkHandleMultipartField",
            "value": 181.6,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "6603757 times\n4 procs"
          },
          {
            "name": "BenchmarkHandleMultipartField - ns/op",
            "value": 181.6,
            "unit": "ns/op",
            "extra": "6603757 times\n4 procs"
          },
          {
            "name": "BenchmarkHandleMultipartField - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "6603757 times\n4 procs"
          },
          {
            "name": "BenchmarkHandleMultipartField - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "6603757 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecode",
            "value": 7069,
            "unit": "ns/op\t    1200 B/op\t      62 allocs/op",
            "extra": "167635 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecode - ns/op",
            "value": 7069,
            "unit": "ns/op",
            "extra": "167635 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecode - B/op",
            "value": 1200,
            "unit": "B/op",
            "extra": "167635 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecode - allocs/op",
            "value": 62,
            "unit": "allocs/op",
            "extra": "167635 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecodeParallel",
            "value": 3415,
            "unit": "ns/op\t    1200 B/op\t      62 allocs/op",
            "extra": "348172 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecodeParallel - ns/op",
            "value": 3415,
            "unit": "ns/op",
            "extra": "348172 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecodeParallel - B/op",
            "value": 1200,
            "unit": "B/op",
            "extra": "348172 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecodeParallel - allocs/op",
            "value": 62,
            "unit": "allocs/op",
            "extra": "348172 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructDecode",
            "value": 3119,
            "unit": "ns/op\t     479 B/op\t      30 allocs/op",
            "extra": "361436 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructDecode - ns/op",
            "value": 3119,
            "unit": "ns/op",
            "extra": "361436 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructDecode - B/op",
            "value": 479,
            "unit": "B/op",
            "extra": "361436 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructDecode - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "361436 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckRequiredFields",
            "value": 1195,
            "unit": "ns/op\t    1048 B/op\t      15 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckRequiredFields - ns/op",
            "value": 1195,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckRequiredFields - B/op",
            "value": 1048,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckRequiredFields - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationDecoding",
            "value": 783.3,
            "unit": "ns/op\t     184 B/op\t       8 allocs/op",
            "extra": "1534647 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationDecoding - ns/op",
            "value": 783.3,
            "unit": "ns/op",
            "extra": "1534647 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationDecoding - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "1534647 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationDecoding - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1534647 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncode",
            "value": 1687,
            "unit": "ns/op\t     575 B/op\t       5 allocs/op",
            "extra": "771777 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncode - ns/op",
            "value": 1687,
            "unit": "ns/op",
            "extra": "771777 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncode - B/op",
            "value": 575,
            "unit": "B/op",
            "extra": "771777 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncode - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "771777 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncodeParallel",
            "value": 760.6,
            "unit": "ns/op\t     577 B/op\t       5 allocs/op",
            "extra": "1479492 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncodeParallel - ns/op",
            "value": 760.6,
            "unit": "ns/op",
            "extra": "1479492 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncodeParallel - B/op",
            "value": 577,
            "unit": "B/op",
            "extra": "1479492 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncodeParallel - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "1479492 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncode",
            "value": 2724,
            "unit": "ns/op\t     993 B/op\t      12 allocs/op",
            "extra": "450254 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncode - ns/op",
            "value": 2724,
            "unit": "ns/op",
            "extra": "450254 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncode - B/op",
            "value": 993,
            "unit": "B/op",
            "extra": "450254 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncode - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "450254 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncodeParallel",
            "value": 1457,
            "unit": "ns/op\t    1049 B/op\t      12 allocs/op",
            "extra": "940072 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncodeParallel - ns/op",
            "value": 1457,
            "unit": "ns/op",
            "extra": "940072 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncodeParallel - B/op",
            "value": 1049,
            "unit": "B/op",
            "extra": "940072 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncodeParallel - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "940072 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationEncoding",
            "value": 344.5,
            "unit": "ns/op\t     147 B/op\t       2 allocs/op",
            "extra": "3452449 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationEncoding - ns/op",
            "value": 344.5,
            "unit": "ns/op",
            "extra": "3452449 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationEncoding - B/op",
            "value": 147,
            "unit": "B/op",
            "extra": "3452449 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationEncoding - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "3452449 times\n4 procs"
          },
          {
            "name": "BenchmarkMultiErrorError",
            "value": 184.5,
            "unit": "ns/op\t      40 B/op\t       2 allocs/op",
            "extra": "6475292 times\n4 procs"
          },
          {
            "name": "BenchmarkMultiErrorError - ns/op",
            "value": 184.5,
            "unit": "ns/op",
            "extra": "6475292 times\n4 procs"
          },
          {
            "name": "BenchmarkMultiErrorError - B/op",
            "value": 40,
            "unit": "B/op",
            "extra": "6475292 times\n4 procs"
          },
          {
            "name": "BenchmarkMultiErrorError - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "6475292 times\n4 procs"
          },
          {
            "name": "BenchmarkParseTag",
            "value": 49.9,
            "unit": "ns/op\t      32 B/op\t       1 allocs/op",
            "extra": "23820207 times\n4 procs"
          },
          {
            "name": "BenchmarkParseTag - ns/op",
            "value": 49.9,
            "unit": "ns/op",
            "extra": "23820207 times\n4 procs"
          },
          {
            "name": "BenchmarkParseTag - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "23820207 times\n4 procs"
          },
          {
            "name": "BenchmarkParseTag - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "23820207 times\n4 procs"
          },
          {
            "name": "BenchmarkIsZero",
            "value": 19.06,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "63463602 times\n4 procs"
          },
          {
            "name": "BenchmarkIsZero - ns/op",
            "value": 19.06,
            "unit": "ns/op",
            "extra": "63463602 times\n4 procs"
          },
          {
            "name": "BenchmarkIsZero - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "63463602 times\n4 procs"
          },
          {
            "name": "BenchmarkIsZero - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "63463602 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertPointer",
            "value": 24.04,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "47386473 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertPointer - ns/op",
            "value": 24.04,
            "unit": "ns/op",
            "extra": "47386473 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertPointer - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "47386473 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertPointer - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "47386473 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "rene@gofiber.io",
            "name": "RW",
            "username": "ReneWerner87"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "ab7e52e7eb179c6f182f940d46695e3a8e80e337",
          "message": "Merge pull request #74 from gofiber/dependabot/go_modules/github.com/gofiber/utils/v2-2.2.0\n\nbuild(deps): bump github.com/gofiber/utils/v2 from 2.1.2 to 2.2.0",
          "timestamp": "2026-07-15T18:03:20+02:00",
          "tree_id": "18b1db172568a749cf92e1aaeedc794c2bcf2b04",
          "url": "https://github.com/gofiber/schema/commit/ab7e52e7eb179c6f182f940d46695e3a8e80e337"
        },
        "date": 1784131654092,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBool",
            "value": 3.459,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "349531867 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBool - ns/op",
            "value": 3.459,
            "unit": "ns/op",
            "extra": "349531867 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBool - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "349531867 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBool - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "349531867 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertInt",
            "value": 8.434,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "142182255 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertInt - ns/op",
            "value": 8.434,
            "unit": "ns/op",
            "extra": "142182255 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertInt - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "142182255 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertInt - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "142182255 times\n4 procs"
          },
          {
            "name": "BenchmarkAll",
            "value": 33.11,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "35694518 times\n4 procs"
          },
          {
            "name": "BenchmarkAll - ns/op",
            "value": 33.11,
            "unit": "ns/op",
            "extra": "35694518 times\n4 procs"
          },
          {
            "name": "BenchmarkAll - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "35694518 times\n4 procs"
          },
          {
            "name": "BenchmarkAll - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "35694518 times\n4 procs"
          },
          {
            "name": "BenchmarkDecoderMultipartFiles",
            "value": 6262,
            "unit": "ns/op\t    1567 B/op\t      48 allocs/op",
            "extra": "185005 times\n4 procs"
          },
          {
            "name": "BenchmarkDecoderMultipartFiles - ns/op",
            "value": 6262,
            "unit": "ns/op",
            "extra": "185005 times\n4 procs"
          },
          {
            "name": "BenchmarkDecoderMultipartFiles - B/op",
            "value": 1567,
            "unit": "B/op",
            "extra": "185005 times\n4 procs"
          },
          {
            "name": "BenchmarkDecoderMultipartFiles - allocs/op",
            "value": 48,
            "unit": "allocs/op",
            "extra": "185005 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-0",
            "value": 6.943,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "171667609 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-0 - ns/op",
            "value": 6.943,
            "unit": "ns/op",
            "extra": "171667609 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-0 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "171667609 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-0 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "171667609 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-1",
            "value": 6.868,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "174660799 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-1 - ns/op",
            "value": 6.868,
            "unit": "ns/op",
            "extra": "174660799 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-1 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "174660799 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-1 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "174660799 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-2",
            "value": 5.913,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "203654289 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-2 - ns/op",
            "value": 5.913,
            "unit": "ns/op",
            "extra": "203654289 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-2 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "203654289 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-2 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "203654289 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-3",
            "value": 6.868,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "174679430 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-3 - ns/op",
            "value": 6.868,
            "unit": "ns/op",
            "extra": "174679430 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-3 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "174679430 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-3 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "174679430 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-4",
            "value": 3.75,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "318698546 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-4 - ns/op",
            "value": 3.75,
            "unit": "ns/op",
            "extra": "318698546 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-4 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "318698546 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-4 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "318698546 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-5",
            "value": 6.95,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "171045837 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-5 - ns/op",
            "value": 6.95,
            "unit": "ns/op",
            "extra": "171045837 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-5 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "171045837 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-5 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "171045837 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-6",
            "value": 12.91,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "96136220 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-6 - ns/op",
            "value": 12.91,
            "unit": "ns/op",
            "extra": "96136220 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-6 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "96136220 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-6 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "96136220 times\n4 procs"
          },
          {
            "name": "BenchmarkHandleMultipartField",
            "value": 182.1,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "6574243 times\n4 procs"
          },
          {
            "name": "BenchmarkHandleMultipartField - ns/op",
            "value": 182.1,
            "unit": "ns/op",
            "extra": "6574243 times\n4 procs"
          },
          {
            "name": "BenchmarkHandleMultipartField - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "6574243 times\n4 procs"
          },
          {
            "name": "BenchmarkHandleMultipartField - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "6574243 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecode",
            "value": 7307,
            "unit": "ns/op\t    1200 B/op\t      62 allocs/op",
            "extra": "161649 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecode - ns/op",
            "value": 7307,
            "unit": "ns/op",
            "extra": "161649 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecode - B/op",
            "value": 1200,
            "unit": "B/op",
            "extra": "161649 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecode - allocs/op",
            "value": 62,
            "unit": "allocs/op",
            "extra": "161649 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecodeParallel",
            "value": 3511,
            "unit": "ns/op\t    1200 B/op\t      62 allocs/op",
            "extra": "338910 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecodeParallel - ns/op",
            "value": 3511,
            "unit": "ns/op",
            "extra": "338910 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecodeParallel - B/op",
            "value": 1200,
            "unit": "B/op",
            "extra": "338910 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecodeParallel - allocs/op",
            "value": 62,
            "unit": "allocs/op",
            "extra": "338910 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructDecode",
            "value": 3252,
            "unit": "ns/op\t     479 B/op\t      30 allocs/op",
            "extra": "350630 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructDecode - ns/op",
            "value": 3252,
            "unit": "ns/op",
            "extra": "350630 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructDecode - B/op",
            "value": 479,
            "unit": "B/op",
            "extra": "350630 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructDecode - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "350630 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckRequiredFields",
            "value": 1196,
            "unit": "ns/op\t    1048 B/op\t      15 allocs/op",
            "extra": "959622 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckRequiredFields - ns/op",
            "value": 1196,
            "unit": "ns/op",
            "extra": "959622 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckRequiredFields - B/op",
            "value": 1048,
            "unit": "B/op",
            "extra": "959622 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckRequiredFields - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "959622 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationDecoding",
            "value": 804.7,
            "unit": "ns/op\t     184 B/op\t       8 allocs/op",
            "extra": "1488164 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationDecoding - ns/op",
            "value": 804.7,
            "unit": "ns/op",
            "extra": "1488164 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationDecoding - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "1488164 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationDecoding - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1488164 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncode",
            "value": 1631,
            "unit": "ns/op\t     538 B/op\t       5 allocs/op",
            "extra": "838831 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncode - ns/op",
            "value": 1631,
            "unit": "ns/op",
            "extra": "838831 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncode - B/op",
            "value": 538,
            "unit": "B/op",
            "extra": "838831 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncode - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "838831 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncodeParallel",
            "value": 775.2,
            "unit": "ns/op\t     604 B/op\t       5 allocs/op",
            "extra": "1470334 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncodeParallel - ns/op",
            "value": 775.2,
            "unit": "ns/op",
            "extra": "1470334 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncodeParallel - B/op",
            "value": 604,
            "unit": "B/op",
            "extra": "1470334 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncodeParallel - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "1470334 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncode",
            "value": 2707,
            "unit": "ns/op\t    1030 B/op\t      12 allocs/op",
            "extra": "425049 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncode - ns/op",
            "value": 2707,
            "unit": "ns/op",
            "extra": "425049 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncode - B/op",
            "value": 1030,
            "unit": "B/op",
            "extra": "425049 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncode - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "425049 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncodeParallel",
            "value": 1479,
            "unit": "ns/op\t    1058 B/op\t      12 allocs/op",
            "extra": "821905 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncodeParallel - ns/op",
            "value": 1479,
            "unit": "ns/op",
            "extra": "821905 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncodeParallel - B/op",
            "value": 1058,
            "unit": "B/op",
            "extra": "821905 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncodeParallel - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "821905 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationEncoding",
            "value": 326.8,
            "unit": "ns/op\t     145 B/op\t       2 allocs/op",
            "extra": "3534559 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationEncoding - ns/op",
            "value": 326.8,
            "unit": "ns/op",
            "extra": "3534559 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationEncoding - B/op",
            "value": 145,
            "unit": "B/op",
            "extra": "3534559 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationEncoding - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "3534559 times\n4 procs"
          },
          {
            "name": "BenchmarkMultiErrorError",
            "value": 177.5,
            "unit": "ns/op\t      40 B/op\t       2 allocs/op",
            "extra": "6750292 times\n4 procs"
          },
          {
            "name": "BenchmarkMultiErrorError - ns/op",
            "value": 177.5,
            "unit": "ns/op",
            "extra": "6750292 times\n4 procs"
          },
          {
            "name": "BenchmarkMultiErrorError - B/op",
            "value": 40,
            "unit": "B/op",
            "extra": "6750292 times\n4 procs"
          },
          {
            "name": "BenchmarkMultiErrorError - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "6750292 times\n4 procs"
          },
          {
            "name": "BenchmarkParseTag",
            "value": 48.39,
            "unit": "ns/op\t      32 B/op\t       1 allocs/op",
            "extra": "24209020 times\n4 procs"
          },
          {
            "name": "BenchmarkParseTag - ns/op",
            "value": 48.39,
            "unit": "ns/op",
            "extra": "24209020 times\n4 procs"
          },
          {
            "name": "BenchmarkParseTag - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "24209020 times\n4 procs"
          },
          {
            "name": "BenchmarkParseTag - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "24209020 times\n4 procs"
          },
          {
            "name": "BenchmarkIsZero",
            "value": 19.03,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "63608250 times\n4 procs"
          },
          {
            "name": "BenchmarkIsZero - ns/op",
            "value": 19.03,
            "unit": "ns/op",
            "extra": "63608250 times\n4 procs"
          },
          {
            "name": "BenchmarkIsZero - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "63608250 times\n4 procs"
          },
          {
            "name": "BenchmarkIsZero - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "63608250 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertPointer",
            "value": 23.37,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "48497418 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertPointer - ns/op",
            "value": 23.37,
            "unit": "ns/op",
            "extra": "48497418 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertPointer - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "48497418 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertPointer - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "48497418 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "rene@gofiber.io",
            "name": "RW",
            "username": "ReneWerner87"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "70537bac89b88769704d2a4acfabbd73dfdda589",
          "message": "Merge pull request #75 from gofiber/dependabot/github_actions/actions/setup-go-7.0.0\n\nbump actions/setup-go from 6.5.0 to 7.0.0",
          "timestamp": "2026-07-17T02:15:26+02:00",
          "tree_id": "187742091314c5dbfb74ffb63ea081547cbe0927",
          "url": "https://github.com/gofiber/schema/commit/70537bac89b88769704d2a4acfabbd73dfdda589"
        },
        "date": 1784247383541,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConvertBool",
            "value": 3.471,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "348226198 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBool - ns/op",
            "value": 3.471,
            "unit": "ns/op",
            "extra": "348226198 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBool - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "348226198 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBool - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "348226198 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertInt",
            "value": 8.486,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "141156267 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertInt - ns/op",
            "value": 8.486,
            "unit": "ns/op",
            "extra": "141156267 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertInt - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "141156267 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertInt - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "141156267 times\n4 procs"
          },
          {
            "name": "BenchmarkAll",
            "value": 33.2,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "36172437 times\n4 procs"
          },
          {
            "name": "BenchmarkAll - ns/op",
            "value": 33.2,
            "unit": "ns/op",
            "extra": "36172437 times\n4 procs"
          },
          {
            "name": "BenchmarkAll - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "36172437 times\n4 procs"
          },
          {
            "name": "BenchmarkAll - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "36172437 times\n4 procs"
          },
          {
            "name": "BenchmarkDecoderMultipartFiles",
            "value": 6217,
            "unit": "ns/op\t    1567 B/op\t      48 allocs/op",
            "extra": "189987 times\n4 procs"
          },
          {
            "name": "BenchmarkDecoderMultipartFiles - ns/op",
            "value": 6217,
            "unit": "ns/op",
            "extra": "189987 times\n4 procs"
          },
          {
            "name": "BenchmarkDecoderMultipartFiles - B/op",
            "value": 1567,
            "unit": "B/op",
            "extra": "189987 times\n4 procs"
          },
          {
            "name": "BenchmarkDecoderMultipartFiles - allocs/op",
            "value": 48,
            "unit": "allocs/op",
            "extra": "189987 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-0",
            "value": 6.869,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "174463178 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-0 - ns/op",
            "value": 6.869,
            "unit": "ns/op",
            "extra": "174463178 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-0 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "174463178 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-0 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "174463178 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-1",
            "value": 6.89,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "173611058 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-1 - ns/op",
            "value": 6.89,
            "unit": "ns/op",
            "extra": "173611058 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-1 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "173611058 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-1 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "173611058 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-2",
            "value": 5.616,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "213658636 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-2 - ns/op",
            "value": 5.616,
            "unit": "ns/op",
            "extra": "213658636 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-2 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "213658636 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-2 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "213658636 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-3",
            "value": 6.856,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "174973563 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-3 - ns/op",
            "value": 6.856,
            "unit": "ns/op",
            "extra": "174973563 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-3 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "174973563 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-3 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "174973563 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-4",
            "value": 3.738,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "321040236 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-4 - ns/op",
            "value": 3.738,
            "unit": "ns/op",
            "extra": "321040236 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-4 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "321040236 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-4 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "321040236 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-5",
            "value": 6.964,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "170889043 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-5 - ns/op",
            "value": 6.964,
            "unit": "ns/op",
            "extra": "170889043 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-5 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "170889043 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-5 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "170889043 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-6",
            "value": 13.18,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "89901169 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-6 - ns/op",
            "value": 13.18,
            "unit": "ns/op",
            "extra": "89901169 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-6 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "89901169 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-6 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "89901169 times\n4 procs"
          },
          {
            "name": "BenchmarkHandleMultipartField",
            "value": 182.9,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "6392036 times\n4 procs"
          },
          {
            "name": "BenchmarkHandleMultipartField - ns/op",
            "value": 182.9,
            "unit": "ns/op",
            "extra": "6392036 times\n4 procs"
          },
          {
            "name": "BenchmarkHandleMultipartField - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "6392036 times\n4 procs"
          },
          {
            "name": "BenchmarkHandleMultipartField - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "6392036 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecode",
            "value": 7014,
            "unit": "ns/op\t    1200 B/op\t      62 allocs/op",
            "extra": "164466 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecode - ns/op",
            "value": 7014,
            "unit": "ns/op",
            "extra": "164466 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecode - B/op",
            "value": 1200,
            "unit": "B/op",
            "extra": "164466 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecode - allocs/op",
            "value": 62,
            "unit": "allocs/op",
            "extra": "164466 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecodeParallel",
            "value": 3495,
            "unit": "ns/op\t    1200 B/op\t      62 allocs/op",
            "extra": "336386 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecodeParallel - ns/op",
            "value": 3495,
            "unit": "ns/op",
            "extra": "336386 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecodeParallel - B/op",
            "value": 1200,
            "unit": "B/op",
            "extra": "336386 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecodeParallel - allocs/op",
            "value": 62,
            "unit": "allocs/op",
            "extra": "336386 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructDecode",
            "value": 3181,
            "unit": "ns/op\t     479 B/op\t      30 allocs/op",
            "extra": "360858 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructDecode - ns/op",
            "value": 3181,
            "unit": "ns/op",
            "extra": "360858 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructDecode - B/op",
            "value": 479,
            "unit": "B/op",
            "extra": "360858 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructDecode - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "360858 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckRequiredFields",
            "value": 1200,
            "unit": "ns/op\t    1048 B/op\t      15 allocs/op",
            "extra": "999649 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckRequiredFields - ns/op",
            "value": 1200,
            "unit": "ns/op",
            "extra": "999649 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckRequiredFields - B/op",
            "value": 1048,
            "unit": "B/op",
            "extra": "999649 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckRequiredFields - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "999649 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationDecoding",
            "value": 783.6,
            "unit": "ns/op\t     184 B/op\t       8 allocs/op",
            "extra": "1530628 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationDecoding - ns/op",
            "value": 783.6,
            "unit": "ns/op",
            "extra": "1530628 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationDecoding - B/op",
            "value": 184,
            "unit": "B/op",
            "extra": "1530628 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationDecoding - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1530628 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncode",
            "value": 1602,
            "unit": "ns/op\t     549 B/op\t       5 allocs/op",
            "extra": "817156 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncode - ns/op",
            "value": 1602,
            "unit": "ns/op",
            "extra": "817156 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncode - B/op",
            "value": 549,
            "unit": "B/op",
            "extra": "817156 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncode - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "817156 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncodeParallel",
            "value": 778.9,
            "unit": "ns/op\t     587 B/op\t       5 allocs/op",
            "extra": "1523089 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncodeParallel - ns/op",
            "value": 778.9,
            "unit": "ns/op",
            "extra": "1523089 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncodeParallel - B/op",
            "value": 587,
            "unit": "B/op",
            "extra": "1523089 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncodeParallel - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "1523089 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncode",
            "value": 2743,
            "unit": "ns/op\t    1060 B/op\t      12 allocs/op",
            "extra": "407068 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncode - ns/op",
            "value": 2743,
            "unit": "ns/op",
            "extra": "407068 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncode - B/op",
            "value": 1060,
            "unit": "B/op",
            "extra": "407068 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncode - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "407068 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncodeParallel",
            "value": 1407,
            "unit": "ns/op\t    1037 B/op\t      12 allocs/op",
            "extra": "902382 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncodeParallel - ns/op",
            "value": 1407,
            "unit": "ns/op",
            "extra": "902382 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncodeParallel - B/op",
            "value": 1037,
            "unit": "B/op",
            "extra": "902382 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncodeParallel - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "902382 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationEncoding",
            "value": 324.1,
            "unit": "ns/op\t     142 B/op\t       2 allocs/op",
            "extra": "3666045 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationEncoding - ns/op",
            "value": 324.1,
            "unit": "ns/op",
            "extra": "3666045 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationEncoding - B/op",
            "value": 142,
            "unit": "B/op",
            "extra": "3666045 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationEncoding - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "3666045 times\n4 procs"
          },
          {
            "name": "BenchmarkMultiErrorError",
            "value": 184.9,
            "unit": "ns/op\t      40 B/op\t       2 allocs/op",
            "extra": "6126271 times\n4 procs"
          },
          {
            "name": "BenchmarkMultiErrorError - ns/op",
            "value": 184.9,
            "unit": "ns/op",
            "extra": "6126271 times\n4 procs"
          },
          {
            "name": "BenchmarkMultiErrorError - B/op",
            "value": 40,
            "unit": "B/op",
            "extra": "6126271 times\n4 procs"
          },
          {
            "name": "BenchmarkMultiErrorError - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "6126271 times\n4 procs"
          },
          {
            "name": "BenchmarkParseTag",
            "value": 49.4,
            "unit": "ns/op\t      32 B/op\t       1 allocs/op",
            "extra": "23974856 times\n4 procs"
          },
          {
            "name": "BenchmarkParseTag - ns/op",
            "value": 49.4,
            "unit": "ns/op",
            "extra": "23974856 times\n4 procs"
          },
          {
            "name": "BenchmarkParseTag - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "23974856 times\n4 procs"
          },
          {
            "name": "BenchmarkParseTag - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "23974856 times\n4 procs"
          },
          {
            "name": "BenchmarkIsZero",
            "value": 19.12,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "63483642 times\n4 procs"
          },
          {
            "name": "BenchmarkIsZero - ns/op",
            "value": 19.12,
            "unit": "ns/op",
            "extra": "63483642 times\n4 procs"
          },
          {
            "name": "BenchmarkIsZero - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "63483642 times\n4 procs"
          },
          {
            "name": "BenchmarkIsZero - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "63483642 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertPointer",
            "value": 23.88,
            "unit": "ns/op\t       8 B/op\t       1 allocs/op",
            "extra": "49937770 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertPointer - ns/op",
            "value": 23.88,
            "unit": "ns/op",
            "extra": "49937770 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertPointer - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "49937770 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertPointer - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "49937770 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "rene@gofiber.io",
            "name": "RW",
            "username": "ReneWerner87"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "9cc21490bc326064a3bef6097e58e3ae28ffc08f",
          "message": "Merge pull request #76 from gofiber/claude/schema-performance-optimization-kw0q51\n\n⚡ perf: rewrite decode/encode hot paths — zero-alloc decoding, SWAR utils, cached metadata, panic-safety",
          "timestamp": "2026-07-19T18:40:32+02:00",
          "tree_id": "b16c25618aa676f3a4217802975c07b194db1c2f",
          "url": "https://github.com/gofiber/schema/commit/9cc21490bc326064a3bef6097e58e3ae28ffc08f"
        },
        "date": 1784479288622,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkNextPathSegment",
            "value": 14.79,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "76902909 times\n4 procs"
          },
          {
            "name": "BenchmarkNextPathSegment - ns/op",
            "value": 14.79,
            "unit": "ns/op",
            "extra": "76902909 times\n4 procs"
          },
          {
            "name": "BenchmarkNextPathSegment - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "76902909 times\n4 procs"
          },
          {
            "name": "BenchmarkNextPathSegment - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "76902909 times\n4 procs"
          },
          {
            "name": "BenchmarkParsePathCacheMiss",
            "value": 463.5,
            "unit": "ns/op\t     504 B/op\t       9 allocs/op",
            "extra": "2792143 times\n4 procs"
          },
          {
            "name": "BenchmarkParsePathCacheMiss - ns/op",
            "value": 463.5,
            "unit": "ns/op",
            "extra": "2792143 times\n4 procs"
          },
          {
            "name": "BenchmarkParsePathCacheMiss - B/op",
            "value": 504,
            "unit": "B/op",
            "extra": "2792143 times\n4 procs"
          },
          {
            "name": "BenchmarkParsePathCacheMiss - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "2792143 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBool",
            "value": 3.536,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "339832544 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBool - ns/op",
            "value": 3.536,
            "unit": "ns/op",
            "extra": "339832544 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBool - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "339832544 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertBool - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "339832544 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertInt",
            "value": 8.512,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "140743347 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertInt - ns/op",
            "value": 8.512,
            "unit": "ns/op",
            "extra": "140743347 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertInt - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "140743347 times\n4 procs"
          },
          {
            "name": "BenchmarkConvertInt - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "140743347 times\n4 procs"
          },
          {
            "name": "BenchmarkAll",
            "value": 3.887,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "308889573 times\n4 procs"
          },
          {
            "name": "BenchmarkAll - ns/op",
            "value": 3.887,
            "unit": "ns/op",
            "extra": "308889573 times\n4 procs"
          },
          {
            "name": "BenchmarkAll - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "308889573 times\n4 procs"
          },
          {
            "name": "BenchmarkAll - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "308889573 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceManyIndicesDecode",
            "value": 35964,
            "unit": "ns/op\t   16517 B/op\t      17 allocs/op",
            "extra": "33145 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceManyIndicesDecode - ns/op",
            "value": 35964,
            "unit": "ns/op",
            "extra": "33145 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceManyIndicesDecode - B/op",
            "value": 16517,
            "unit": "B/op",
            "extra": "33145 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceManyIndicesDecode - allocs/op",
            "value": 17,
            "unit": "allocs/op",
            "extra": "33145 times\n4 procs"
          },
          {
            "name": "BenchmarkDecoderMultipartFiles",
            "value": 1716,
            "unit": "ns/op\t      96 B/op\t       5 allocs/op",
            "extra": "675006 times\n4 procs"
          },
          {
            "name": "BenchmarkDecoderMultipartFiles - ns/op",
            "value": 1716,
            "unit": "ns/op",
            "extra": "675006 times\n4 procs"
          },
          {
            "name": "BenchmarkDecoderMultipartFiles - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "675006 times\n4 procs"
          },
          {
            "name": "BenchmarkDecoderMultipartFiles - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "675006 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-0",
            "value": 8.064,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "148608501 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-0 - ns/op",
            "value": 8.064,
            "unit": "ns/op",
            "extra": "148608501 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-0 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "148608501 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-0 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "148608501 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-1",
            "value": 8.018,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "150544982 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-1 - ns/op",
            "value": 8.018,
            "unit": "ns/op",
            "extra": "150544982 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-1 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "150544982 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-1 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "150544982 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-2",
            "value": 6.336,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "188301770 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-2 - ns/op",
            "value": 6.336,
            "unit": "ns/op",
            "extra": "188301770 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-2 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "188301770 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-2 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "188301770 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-3",
            "value": 8.104,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "148029424 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-3 - ns/op",
            "value": 8.104,
            "unit": "ns/op",
            "extra": "148029424 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-3 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "148029424 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-3 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "148029424 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-4",
            "value": 4.082,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "284595434 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-4 - ns/op",
            "value": 4.082,
            "unit": "ns/op",
            "extra": "284595434 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-4 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "284595434 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-4 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "284595434 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-5",
            "value": 8.188,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "148378770 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-5 - ns/op",
            "value": 8.188,
            "unit": "ns/op",
            "extra": "148378770 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-5 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "148378770 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-5 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "148378770 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-6",
            "value": 15.3,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "80768126 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-6 - ns/op",
            "value": 15.3,
            "unit": "ns/op",
            "extra": "80768126 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-6 - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "80768126 times\n4 procs"
          },
          {
            "name": "BenchmarkIsMultipartFile/IsMultipartFile-6 - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "80768126 times\n4 procs"
          },
          {
            "name": "BenchmarkHandleMultipartField",
            "value": 175.2,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "6848871 times\n4 procs"
          },
          {
            "name": "BenchmarkHandleMultipartField - ns/op",
            "value": 175.2,
            "unit": "ns/op",
            "extra": "6848871 times\n4 procs"
          },
          {
            "name": "BenchmarkHandleMultipartField - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "6848871 times\n4 procs"
          },
          {
            "name": "BenchmarkHandleMultipartField - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "6848871 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecode",
            "value": 1369,
            "unit": "ns/op\t     144 B/op\t       4 allocs/op",
            "extra": "833618 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecode - ns/op",
            "value": 1369,
            "unit": "ns/op",
            "extra": "833618 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecode - B/op",
            "value": 144,
            "unit": "B/op",
            "extra": "833618 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecode - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "833618 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecodeParallel",
            "value": 524.2,
            "unit": "ns/op\t     144 B/op\t       4 allocs/op",
            "extra": "2301181 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecodeParallel - ns/op",
            "value": 524.2,
            "unit": "ns/op",
            "extra": "2301181 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecodeParallel - B/op",
            "value": 144,
            "unit": "B/op",
            "extra": "2301181 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructDecodeParallel - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "2301181 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructDecode",
            "value": 596.6,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "2016276 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructDecode - ns/op",
            "value": 596.6,
            "unit": "ns/op",
            "extra": "2016276 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructDecode - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "2016276 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructDecode - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "2016276 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckRequiredFields",
            "value": 235.9,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "4912492 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckRequiredFields - ns/op",
            "value": 235.9,
            "unit": "ns/op",
            "extra": "4912492 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckRequiredFields - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "4912492 times\n4 procs"
          },
          {
            "name": "BenchmarkCheckRequiredFields - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "4912492 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationDecoding",
            "value": 273,
            "unit": "ns/op\t      16 B/op\t       2 allocs/op",
            "extra": "4361821 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationDecoding - ns/op",
            "value": 273,
            "unit": "ns/op",
            "extra": "4361821 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationDecoding - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "4361821 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationDecoding - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "4361821 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncode",
            "value": 948,
            "unit": "ns/op\t     462 B/op\t       3 allocs/op",
            "extra": "1257930 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncode - ns/op",
            "value": 948,
            "unit": "ns/op",
            "extra": "1257930 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncode - B/op",
            "value": 462,
            "unit": "B/op",
            "extra": "1257930 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncode - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "1257930 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncodeParallel",
            "value": 416.8,
            "unit": "ns/op\t     456 B/op\t       3 allocs/op",
            "extra": "2758026 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncodeParallel - ns/op",
            "value": 416.8,
            "unit": "ns/op",
            "extra": "2758026 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncodeParallel - B/op",
            "value": 456,
            "unit": "B/op",
            "extra": "2758026 times\n4 procs"
          },
          {
            "name": "BenchmarkSimpleStructEncodeParallel - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "2758026 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncode",
            "value": 1116,
            "unit": "ns/op\t     872 B/op\t       6 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncode - ns/op",
            "value": 1116,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncode - B/op",
            "value": 872,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncode - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncodeParallel",
            "value": 662.1,
            "unit": "ns/op\t     924 B/op\t       6 allocs/op",
            "extra": "1891161 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncodeParallel - ns/op",
            "value": 662.1,
            "unit": "ns/op",
            "extra": "1891161 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncodeParallel - B/op",
            "value": 924,
            "unit": "B/op",
            "extra": "1891161 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeStructEncodeParallel - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1891161 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationEncoding",
            "value": 130.9,
            "unit": "ns/op\t      85 B/op\t       1 allocs/op",
            "extra": "9892609 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationEncoding - ns/op",
            "value": 130.9,
            "unit": "ns/op",
            "extra": "9892609 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationEncoding - B/op",
            "value": 85,
            "unit": "B/op",
            "extra": "9892609 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeDurationEncoding - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "9892609 times\n4 procs"
          },
          {
            "name": "BenchmarkMultiErrorError",
            "value": 171.1,
            "unit": "ns/op\t      40 B/op\t       2 allocs/op",
            "extra": "6977122 times\n4 procs"
          },
          {
            "name": "BenchmarkMultiErrorError - ns/op",
            "value": 171.1,
            "unit": "ns/op",
            "extra": "6977122 times\n4 procs"
          },
          {
            "name": "BenchmarkMultiErrorError - B/op",
            "value": 40,
            "unit": "B/op",
            "extra": "6977122 times\n4 procs"
          },
          {
            "name": "BenchmarkMultiErrorError - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "6977122 times\n4 procs"
          },
          {
            "name": "BenchmarkParseTag",
            "value": 4.949,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "242330968 times\n4 procs"
          },
          {
            "name": "BenchmarkParseTag - ns/op",
            "value": 4.949,
            "unit": "ns/op",
            "extra": "242330968 times\n4 procs"
          },
          {
            "name": "BenchmarkParseTag - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "242330968 times\n4 procs"
          },
          {
            "name": "BenchmarkParseTag - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "242330968 times\n4 procs"
          },
          {
            "name": "BenchmarkIsZero",
            "value": 20.43,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "59047428 times\n4 procs"
          },
          {
            "name": "BenchmarkIsZero - ns/op",
            "value": 20.43,
            "unit": "ns/op",
            "extra": "59047428 times\n4 procs"
          },
          {
            "name": "BenchmarkIsZero - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "59047428 times\n4 procs"
          },
          {
            "name": "BenchmarkIsZero - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "59047428 times\n4 procs"
          }
        ]
      }
    ]
  }
}