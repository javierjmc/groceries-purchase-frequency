# Groceries Purchasing Frequency

The purpose of this code is to get statistics on the purchase frequency for groceries, which in theory could be applied to anything you want to get "occurrence frecuency" for.

Inspired on https://blog.smile.io/how-to-calculate-purchase-frequency.

To learn the background behind this code, please refer to my [blog post](https://dev.to/javiermendonca/how-i-hacked-the-way-i-do-groceries-1m0a).

## Preparing the data

- Place your source input as csv files in the `sources` folder. They should be named with the format YYYY-MM-dd.csv, as the name is used to calculate the timeframe in which the calculations will be computed.

- Normalize the data, so you get good results.

- customize the right `path` for the source files if you don't use the given `sources` folder.

run the code as `go run path/to/file/orders.go`
