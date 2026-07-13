window.BENCHMARK_DATA = {
  "lastUpdate": 1783939910196,
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
      }
    ]
  }
}