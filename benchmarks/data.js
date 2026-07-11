window.BENCHMARK_DATA = {
  "lastUpdate": 1783807047096,
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
      }
    ]
  }
}